package usecase

import (
	http "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	"github.com/Miroslovelife/whareflow/internal/services"
	"reflect"
)

type AuthUsecase interface {
	Auth(token, secret string) (bool, string, error)
	Refresh(secret string, expiry int, in *http.JWTCustomClaims) (string, error)
}

type IAuthUsecase struct {
	tokenManager services.TokenManager
}

func NewIAuthUsecase(tokenManager services.TokenManager) *IAuthUsecase {
	return &IAuthUsecase{
		tokenManager: tokenManager,
	}
}

func (au *IAuthUsecase) Auth(token, secret string) (bool, string, error) {
	if ok, err := au.tokenManager.IsAuthorized(token, secret); !ok {
		if err != nil {
			return false, "", err
		}

		return false, "", err
	}

	userId, err := au.tokenManager.ExtractUuidFromToken(token, secret)
	if err != nil {
		return false, "", err
	}

	return true, userId, nil
}

func (au *IAuthUsecase) Refresh(secret string, expiry int, in *http.JWTCustomClaims) (string, error) {

	claimsArgs := make(map[string]interface{})
	val := reflect.ValueOf(in)

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		value := val.Field(i)

		claimsArgs[field.Name] = value.Interface()
	}

	token, err := au.tokenManager.CreateToken(secret, expiry, claimsArgs)
	if err != nil {
		return "", err
	}

	return token, nil
}
