package usecase

import (
	"github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	"github.com/Miroslovelife/whareflow/internal/domain"
	"github.com/Miroslovelife/whareflow/internal/errors"
	"github.com/Miroslovelife/whareflow/internal/repositories"
	"github.com/Miroslovelife/whareflow/internal/services"
)

type UserUsecase interface {
	Register(in *delivery.UserReg) error
	LoginByEmail(in *delivery.UserLoginByEmail, secretAccess string, secretRefresh string, expiry uint8) (string, string, error)
	LoginByPhoneNumber(in *delivery.UserLoginByPhoneNumber, secretAccess string, secretRefresh string, expiry uint8) (string, string, error)
	Refresh(refreshToken, secretAccess, secretRefresh string, expAccess, expRefresh uint8) (string, string, error)
	IsAdmin(userId string) (bool, error)
	IsEmployer(userId string) (bool, error)
}

type userUsecaseImpl struct {
	userRepository repositories.UserRepository
	passwordHasher services.PasswordHasher
	tokenManager   services.TokenManager
}

func NewUserUsecase(userRepository repositories.UserRepository, passwordHasher services.PasswordHasher, tokenManager services.TokenManager) *userUsecaseImpl {
	return &userUsecaseImpl{
		userRepository: userRepository,
		passwordHasher: passwordHasher,
		tokenManager:   tokenManager,
	}
}

func (us *userUsecaseImpl) Register(in *delivery.UserReg) error {

	hashedPassword := us.passwordHasher.Hash(in.Password)

	insertUserData := &domain.User{
		PhoneNumber: in.PhoneNumber,
		Username:    in.Username,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		Surname:     in.Surname,
		Email:       in.Email,
		Password:    hashedPassword,
		Role:        "user",
	}

	if err := us.userRepository.InsertUserData(insertUserData); err != nil {
		return err
	}

	return nil
}

func (us *userUsecaseImpl) LoginByEmail(in *delivery.UserLoginByEmail, secretAccess string, secretRefresh string, expiry uint8) (string, string, error) {
	hashedPassword := us.passwordHasher.Hash(in.Password)

	loginData := map[string]interface{}{
		"email":    in.Email,
		"password": hashedPassword,
	}

	userExist, err := us.userRepository.FindUserData(loginData)
	if err != nil {
		return "", "", err
	}

	claimsAccess := map[string]interface{}{
		"userId":   string(userExist.Uuid),
		"username": userExist.Username,
	}

	claimsRefresh := map[string]interface{}{
		"userId": userExist.Uuid,
	}

	accessToken, err := us.tokenManager.CreateToken(secretAccess, expiry, claimsAccess)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := us.tokenManager.CreateToken(secretRefresh, expiry, claimsRefresh)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil

}

func (us *userUsecaseImpl) LoginByPhoneNumber(in *delivery.UserLoginByPhoneNumber, secretAccess string, secretRefresh string, expiry uint8) (string, string, error) {
	hashedPassword := us.passwordHasher.Hash(in.Password)

	loginData := map[string]interface{}{
		"phone_number": in.PhoneNumber,
		"password":     hashedPassword,
	}

	userExist, err := us.userRepository.FindUserData(loginData)
	if err != nil {
		return "", "", err
	}

	claimsAccess := map[string]interface{}{
		"userId":   string(userExist.Uuid),
		"username": userExist.Username,
	}

	claimsRefresh := map[string]interface{}{
		"userId": userExist.Uuid,
	}

	accessToken, err := us.tokenManager.CreateToken(secretAccess, expiry, claimsAccess)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := us.tokenManager.CreateToken(secretRefresh, expiry, claimsRefresh)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (us *userUsecaseImpl) Refresh(refreshToken, secretAccess, secretRefresh string, expAccess, expRefresh uint8) (string, string, error) {
	if auth, err := us.tokenManager.IsAuthorized(refreshToken, secretRefresh); !auth {
		return "", "", errors.ErrTokenIsNotValid
	} else if err != nil {
		return "", "", err
	}

	userId, err := us.tokenManager.ExtractUuidFromToken(refreshToken, secretRefresh)
	if err != nil {
		return "", "", err
	}

	claims := map[string]interface{}{
		"uuid": userId,
	}

	newAccessToken, err := us.tokenManager.CreateToken(secretAccess, expAccess, claims)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := us.tokenManager.CreateToken(secretRefresh, expRefresh, claims)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

func (us *userUsecaseImpl) IsAdmin(userId string) (bool, error) {
	user, err := us.userRepository.FindUserData(map[string]interface{}{
		"uuid": userId,
	})
	if err != nil {
		return false, err
	}

	if user.Role != "admin" {
		return false, nil
	}

	return true, nil
}

func (us *userUsecaseImpl) IsEmployer(userId string) (bool, error) {
	user, err := us.userRepository.FindUserData(map[string]interface{}{
		"uuid": userId,
	})
	if err != nil {
		return false, err
	}

	if user.Role != "employer" {
		return false, nil
	}

	return true, nil
}
