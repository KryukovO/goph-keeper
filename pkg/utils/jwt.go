package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	// ErrUnexpectedSigningMethod возвращается, если метод подписи JWT
	// не соответствует HMAC-SHA.
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	// ErrTokenIsInvalid возвращается если JWT невалиден.
	ErrTokenIsInvalid = errors.New("token is invalid")
)

// Claims описывает полезную нагрузку JWT.
type Claims struct {
	jwt.RegisteredClaims
	Payload interface{}
}

// BuildJWTString возвращает строковое представление JWT, генерируемого с использованием secret,
// временем жизни lifetime и содержащим payload.
func BuildJWTString(secret []byte, lifetime time.Duration, payload interface{}) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		Claims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(lifetime)),
			},
			Payload: payload,
		},
	)

	return token.SignedString(secret)
}

// ParseTokenString выполняет валидацию JWT.
func ParseTokenString(dst interface{}, tokenString string, secret []byte) error {
	claims := &Claims{
		Payload: dst,
	}

	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("%w: %v", ErrUnexpectedSigningMethod, t.Header["alg"])
			}

			return secret, nil
		},
	)
	if err != nil {
		return err
	}

	if !token.Valid {
		return ErrTokenIsInvalid
	}

	return nil
}
