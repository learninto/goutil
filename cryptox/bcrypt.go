package cryptox

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

//type Bcrypt interface {
//	GenerateFromPassword(ctx context.Context, password string, cost int) (string, error)
//	PasswordVerify(ctx context.Context, password, hash string) error
//}

func GenerateFromPassword(ctx context.Context, password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func CompareHashAndPassword(ctx context.Context, hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
