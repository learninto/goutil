package jwt

import (
	"encoding/json"
	"time"

	"github.com/learninto/goutil/conf"
	"github.com/learninto/goutil/errors"

	jwtD "github.com/dgrijalva/jwt-go"
)

// 一些常量
var (
	SignKey     string = "gwt_sign_key"
	ExpiresTime int    = 259200
)

// Structured version of Claims Section, as referenced at
// https://tools.ietf.org/html/rfc7519#section-4.1
// See examples for how to use this with your own claim types
type CustomClaims struct {
	Data json.RawMessage
	jwtD.StandardClaims
}

// JWT 签名结构
type JWT struct {
	SigningKey  []byte
	ExpiresTime int64
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	signingKey := getSignKey()
	expiresTime := getExpiresTime()
	return &JWT{SigningKey: signingKey, ExpiresTime: expiresTime}
}

// getSignKey 获取signingKey
func getSignKey() []byte {
	signingKey := conf.Get("JWT_SIGNING_KEY")
	if len(signingKey) == 0 {
		signingKey = SignKey // 默认signing key
	}

	return []byte(signingKey)
}

//getExpiresTime 获取过期时间
func getExpiresTime() int64 {
	jTime := conf.GetInt("JWT_EFFECTIVE_DURATION")
	if jTime == 0 {
		jTime = ExpiresTime
	}

	return time.Now().Add(time.Duration(jTime) * time.Second).Unix()
}

// CreateToken 生成一个token
func (j JWT) CreateToken(claims CustomClaims) (string, error) {
	var method jwtD.SigningMethod
	switch conf.Get("JWT_SIGNING_METHOD") {
	case "HS256":
		method = jwtD.SigningMethodHS256
	case "HS384":
		method = jwtD.SigningMethodHS384
	case "HS512":
		method = jwtD.SigningMethodHS512
	default:
		method = jwtD.SigningMethodHS256
	}

	claims.StandardClaims.ExpiresAt = j.ExpiresTime
	token := jwtD.NewWithClaims(method, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Token
func (j JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	if tokenString == "" {
		return nil, errors.TokenMalformed
	}

	token, err := jwtD.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtD.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwtD.ValidationError); ok {
			if ve.Errors&jwtD.ValidationErrorMalformed != 0 {
				return nil, errors.TokenMalformed
			} else if ve.Errors&jwtD.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errors.TokenExpired
			} else if ve.Errors&jwtD.ValidationErrorNotValidYet != 0 {
				return nil, errors.TokenNotValidYet
			} else {
				return nil, errors.TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.TokenInvalid
}

// 更新token
func (j JWT) RefreshToken(tokenString string) (string, error) {
	jwtD.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwtD.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtD.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwtD.TimeFunc = time.Now
		return j.CreateToken(*claims)
	}
	return "", errors.TokenInvalid
}
