package jwtx

import (
	"context"
	"encoding/json"
	"time"

	"github.com/learninto/goutil/conf"
	"github.com/learninto/goutil/errors"

	"github.com/dgrijalva/jwt-go"
)

// 一些常量
var (
	SignKey     string = "gwt_sign_key"
	ExpiresTime int    = 259200
)

// GetSigningKey 获取signingKey
func GetSigningKey() []byte {
	signingKey := conf.Get("JWT_SIGNING_KEY")
	if len(signingKey) == 0 {
		signingKey = SignKey // 默认signing key
	}

	return []byte(signingKey)
}

// GetExpiresTime 获取过期时间
func GetExpiresTime() int64 {
	jTime := conf.GetInt("JWT_EFFECTIVE_DURATION")
	if jTime == 0 {
		jTime = ExpiresTime
	}

	return time.Now().Add(time.Duration(jTime) * time.Second).Unix()
}

// GetSigningMethod 获取签名方法
func GetSigningMethod() (method jwt.SigningMethod) {
	switch conf.Get("JWT_SIGNING_METHOD") {
	case "HS256":
		method = jwt.SigningMethodHS256
	case "HS384":
		method = jwt.SigningMethodHS384
	case "HS512":
		method = jwt.SigningMethodHS512
	default:
		method = jwt.SigningMethodHS256
	}
	return method
}

// CustomClaims
// Structured version of Claims Section, as referenced at
// https://tools.ietf.org/html/rfc7519#section-4.1
// See examples for how to use this with your own claim types
type CustomClaims struct {
	Data json.RawMessage
	jwt.StandardClaims
}

// CreateToken 生成一个token
func (claims CustomClaims) CreateToken(ctx context.Context) (string, error) {
	expiresTime := GetExpiresTime()
	claims.StandardClaims.ExpiresAt = expiresTime

	method := GetSigningMethod()
	token := jwt.NewWithClaims(method, claims)

	signingKey := GetSigningKey()
	return token.SignedString(signingKey)
}

// ParseToken 解析Token
func (CustomClaims) ParseToken(ctx context.Context, tokenString string) (*CustomClaims, error) {
	if tokenString == "" {
		return nil, errors.TokenMalformed
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return GetSigningKey(), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errors.TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
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

// RefreshToken 更新token
func (claims *CustomClaims) RefreshToken(ctx context.Context, tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return GetSigningKey(), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		return claims.CreateToken(ctx)
	}
	return "", errors.TokenInvalid
}
