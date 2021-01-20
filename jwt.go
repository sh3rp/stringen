package stringen

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWTToken(secret string, claims map[string]interface{}) string {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims(claims),
	)

	str, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Printf("Error generating JWT token: %+v", err)
		return ""
	}

	return str
}

func DecodeJWTToken(secret string, tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %+v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		bytes, err := json.Marshal(claims)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}
	return "", fmt.Errorf("Error getting claims")
}
