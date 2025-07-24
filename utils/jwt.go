package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecretkey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := parseToken(token)
	if err := validateToken(parsedToken, err); err != nil {
		return 0, err
	}
	// If you want to use claims:
	claims, err := extractClaims(parsedToken)
	if err != nil {
		return 0, err
	}
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil
}

func parseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
}

func validateToken(parsedToken *jwt.Token, err error) error {
	if err != nil {
		return errors.New("Could not parse token")
	}
	if !parsedToken.Valid {
		return errors.New("Invalid token")
	}
	return nil
}

func extractClaims(parsedToken *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid token claims")
	}
	return claims, nil
}