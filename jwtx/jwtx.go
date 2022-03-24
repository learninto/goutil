package jwtx

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/learninto/goutil/conf"
	"github.com/learninto/goutil/errors"

	"github.com/golang-jwt/jwt/v4"
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

/*
	对应java版本
	// @param privateKey 私钥
	// @param payload 负载，可为空
	// @param issuer jwt颁发者标识
	// @param expireAt jwt过期时间
	PrivateKey priKey = KeyFactory.
		getInstance("RSA").
		generatePrivate(
			new PKCS8EncodedKeySpec(
				(new BASE64Decoder()).decodeBuffer(privateKey)
			)
		);
	Algorithm algorithm = Algorithm.RSA256((RSAPublicKey)null, (RSAPrivateKey)priKey);
	return JWT.create().withIssuer(issuer).withExpiresAt(expireAt).withPayload(payload).sign(algorithm);
*/
func ParsePKCS8PrivateKeyJwt(privateKey string) (tokenStr string, err error) {
	base64DecodeBytes, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return
	}

	pkcs8, err := x509.ParsePKCS8PrivateKey(base64DecodeBytes)
	if err != nil {
		return
	}
	key := pkcs8.(*rsa.PrivateKey)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		Issuer:    "dXzbhbQHY5qnejHa",
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(300 * time.Second)),
	})

	if tokenStr, err = token.SignedString(key); err != nil {
		return
	}
	return
}
