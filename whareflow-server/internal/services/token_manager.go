package services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenManager interface {
	CreateToken(secret string, expiry uint8, args map[string]interface{}) (string, error)
	IsAuthorized(requestToken, secret string) (bool, error)
	ExtractUuidFromToken(requestToken, secret string) (string, error)
}

type TokenM struct {
}

func NewTokenM() *TokenM {
	return &TokenM{}
}

type DynamicCustomClaims struct {
	jwt.RegisteredClaims
	CustomClaims map[string]interface{}
}

func (ja *TokenM) CreateToken(secret string, expiry uint8, args map[string]interface{}) (string, error) {

	exp := jwt.NumericDate{
		Time: time.Now().Add(time.Hour * time.Duration(expiry)),
	}

	claims := jwt.MapClaims{
		"exp": exp,
	}

	for key, value := range args {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (ja *TokenM) IsAuthorized(requestToken, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (ja *TokenM) ExtractUuidFromToken(requestToken, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return "", fmt.Errorf("invalid Token")
	}
	return claims["userId"].(string), nil
}
