package util

import (
	"time"

	"github.com/arvinpaundra/technical-test-api-mnc/config"
	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	UserId string `json:"user_id,omitempty"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId string, duration time.Duration) (string, error) {
	claims := JWTCustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now().Local()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration).Local()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.GetJWTSecret()))
}

func DecodeJWT(tokenStr string) (JWTCustomClaims, error) {
	token, err := jwt.Parse(tokenStr, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenMalformed
		}

		return []byte(config.GetJWTSecret()), nil
	})

	if err != nil {
		return JWTCustomClaims{}, err
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)

		userId := claims["user_id"].(string)

		customClaims := JWTCustomClaims{
			UserId: userId,
		}

		return customClaims, nil
	}

	return JWTCustomClaims{}, jwt.ErrTokenUnverifiable
}
