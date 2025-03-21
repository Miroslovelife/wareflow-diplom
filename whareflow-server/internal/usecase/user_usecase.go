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
	LoginByEmail(in *delivery.UserLoginByEmail, secretAccess string, secretRefresh string, expiry int) (string, string, error)
	LoginByPhoneNumber(in *delivery.UserLoginByPhoneNumber, secretAccess string, secretRefresh string, expiry int) (string, string, error)
	Refresh(refreshToken, secretAccess, secretRefresh string, expAccess, expRefresh int) (string, string, error)
	IsAdmin(userId string) (bool, error)
	IsOwner(userId string) (bool, error)
	IsEmployer(userId string) (bool, error)
	GetProfile(userId string) (*delivery.UserReg, error)
}

type IUserUsecase struct {
	userRepository repositories.UserRepository
	passwordHasher services.PasswordHasher
	tokenManager   services.TokenManager
}

func NewUserUsecase(userRepository repositories.UserRepository, passwordHasher services.PasswordHasher, tokenManager services.TokenManager) *IUserUsecase {
	return &IUserUsecase{
		userRepository: userRepository,
		passwordHasher: passwordHasher,
		tokenManager:   tokenManager,
	}
}

func (us *IUserUsecase) Register(in *delivery.UserReg) error {

	hashedPassword := us.passwordHasher.Hash(in.Password)

	insertUserData := &domain.User{
		PhoneNumber: in.PhoneNumber,
		Username:    in.Username,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		Surname:     in.Surname,
		Email:       in.Email,
		Password:    hashedPassword,
		Role:        in.Role,
	}

	if err := us.userRepository.InsertUserData(insertUserData); err != nil {
		return err
	}

	return nil
}

func (us *IUserUsecase) LoginByEmail(in *delivery.UserLoginByEmail, secretAccess string, secretRefresh string, expiry int) (string, string, error) {
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
		"username": userExist.Username,
		"role":     userExist.Role,
	}

	claimsRefresh := map[string]interface{}{
		"username": userExist.Username,
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

func (us *IUserUsecase) LoginByPhoneNumber(in *delivery.UserLoginByPhoneNumber, secretAccess string, secretRefresh string, expiry int) (string, string, error) {
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
		"username": userExist.Username,
		"role":     userExist.Role,
	}

	claimsRefresh := map[string]interface{}{
		"username": userExist.Username,
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

func (us *IUserUsecase) Refresh(refreshToken, secretAccess, secretRefresh string, expAccess, expRefresh int) (string, string, error) {
	if auth, err := us.tokenManager.IsAuthorized(refreshToken, secretRefresh); !auth {
		return "", "", errors.ErrTokenIsNotValid
	} else if err != nil {
		return "", "", err
	}

	username, err := us.tokenManager.ExtractUsernameToken(refreshToken, secretRefresh)
	if err != nil {
		return "", "", err
	}

	userData := map[string]interface{}{
		"username": username,
	}

	userExist, err := us.userRepository.FindUserData(userData)
	if err != nil {
		return "", "", err
	}

	claimsAccess := map[string]interface{}{
		"username": userExist.Username,
		"role":     userExist.Role,
	}

	claimsRefresh := map[string]interface{}{
		"username": userExist.Username,
	}

	newAccessToken, err := us.tokenManager.CreateToken(secretAccess, expAccess, claimsAccess)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := us.tokenManager.CreateToken(secretRefresh, expRefresh, claimsRefresh)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

func (us *IUserUsecase) IsOwner(userId string) (bool, error) {
	user, err := us.userRepository.FindUserData(map[string]interface{}{
		"uuid": userId,
	})
	if err != nil {
		return false, err
	}

	if user.Role != "owner" {
		return false, nil
	}

	return true, nil
}

func (us *IUserUsecase) IsAdmin(userId string) (bool, error) {
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

func (us *IUserUsecase) IsEmployer(userId string) (bool, error) {
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

func (us *IUserUsecase) GetProfile(userId string) (*delivery.UserReg, error) {
	var profile delivery.UserReg

	data := map[string]interface{}{
		"uuid": userId,
	}

	profileRepo, err := us.userRepository.FindUserData(data)
	if err != nil {
		return nil, err
	}

	profile = delivery.UserReg{
		PhoneNumber: profileRepo.PhoneNumber,
		Username:    profileRepo.Username,
		FirstName:   profileRepo.FirstName,
		LastName:    profileRepo.LastName,
		Surname:     profileRepo.Surname,
		Email:       profileRepo.Email,
		Role:        profileRepo.Role,
	}

	return &profile, nil
}
