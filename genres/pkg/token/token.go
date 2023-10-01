package token

import (
	"github.com/golang-jwt/jwt/v4"
)

// ParseJWT - decode JWT token or return error
func ParseJWT(token string, key []byte) (*Claims, error) {

	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	}

	return &Claims{}, err
}
