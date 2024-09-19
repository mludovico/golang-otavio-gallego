package authentication

import (
	"devbook_api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.SecretKey)
}

func ValidateToken(r *http.Request) error {
	token := extractToken(r)
	parsedToken, err := jwt.Parse(token, getValidationKey)
	if err != nil {
		return err
	}
	if _, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return nil
	}
	return errors.New("invalid token")
}

func ExtractUserId(r *http.Request) (uint64, error) {
	token := extractToken(r)
	parsedToken, err := jwt.Parse(token, getValidationKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return userId, nil
	}
	return 0, errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func getValidationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("invalid signing method: %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
