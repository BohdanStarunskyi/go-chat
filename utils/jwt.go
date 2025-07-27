package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func getSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("secret is empty")
	}
	return secret
}

func GetJwtToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(getSecret()))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateJwt(tokenStr string) (int64, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(getSecret()), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); !ok || int64(exp) < time.Now().Unix() {
			return 0, errors.New("token expired")
		}

		if uid, ok := claims["userID"].(float64); ok {
			return int64(uid), nil
		}
		return 0, errors.New("userID claim not found or invalid")
	}

	return 0, errors.New("invalid token")
}
