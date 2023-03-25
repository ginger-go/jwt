package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// FromToken parses a token and returns the claims.
func FromToken(token, secret string) (*Claims, error) {
	jwtClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	mapClaims, ok := jwtClaims.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}
	claims := NewClaims()
	for key, value := range mapClaims {
		claims.Set(key, value)
	}

	return claims, nil
}

// ToToken creates a token from the claims.
func ToToken(claims *Claims, secret string, expiredAfter time.Duration) (string, error) {
	jwtClaims := jwt.MapClaims{}
	jwtClaims["exp"] = time.Now().Add(expiredAfter).Unix()
	for key, value := range claims.data {
		jwtClaims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return token.SignedString([]byte(secret))
}
