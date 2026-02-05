package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET = []byte("super_secret_key_ganti_env") // nanti pindah ke env

func GenerateToken(userID uint, email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 1 day
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWT_SECRET)
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET, nil
	})
	if err != nil {
		return nil, err
	}

	if BlacklistContains(tokenStr) {
		return token, errors.New("token is blacklisted")
	}

	return token, nil
}

func LogoutToken(tokenStr string) error {
	if tokenStr == "" {
		return nil
	}

	t, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET, nil
	})
	if err != nil {
		// If token can't be parsed, still return nil so logout is idempotent
		return nil
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}

	// extract exp claim
	var expUnix int64
	if expVal, ok := claims["exp"]; ok {
		switch v := expVal.(type) {
		case float64:
			expUnix = int64(v)
		case int64:
			expUnix = v
		}
	}

	// if expUnix is zero, set a conservative expiry (1 hour)
	if expUnix == 0 {
		expUnix = time.Now().Add(time.Hour).Unix()
	}

	BlacklistAdd(tokenStr, expUnix)
	return nil
}
