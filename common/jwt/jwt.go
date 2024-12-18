package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secretKey = "your-secret-key"

func Generate(data map[string]interface{}, expirationTime time.Duration) (string, error) {
	claims := jwt.MapClaims{}

	for key, value := range data {
		claims[key] = value
	}

	claims["exp"] = time.Now().Add(expirationTime).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func Parse(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
