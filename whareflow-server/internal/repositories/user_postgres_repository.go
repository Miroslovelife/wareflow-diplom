package repositories

import (
	"errors"
	"fmt"
	"github.com/Miroslovelife/whareflow/internal/domain"
	error_custom "github.com/Miroslovelife/whareflow/internal/errors"
	"github.com/Miroslovelife/whareflow/pkg/database"
	"gorm.io/gorm"
	"log/slog"
)

type UserRepository interface {
	InsertUserData(in *domain.User) error
	UpdateUserData(in *domain.User) error
	FindUserData(filter map[string]interface{}) (*domain.User, error)
	DeleteUserData(uuid string) error
}

type UserPostgresRepository struct {
	db     database.Database
	logger slog.Logger
}

func NewUserPostgresRepository(db database.Database, logger slog.Logger) *UserPostgresRepository {
	return &UserPostgresRepository{
		db:     db,
		logger: logger,
	}
}

func (ur *UserPostgresRepository) InsertUserData(in *domain.User) error {
	data := &domain.User{
		PhoneNumber: in.PhoneNumber,
		Username:    in.Username,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		Surname:     in.Surname,
		Email:       in.Email,
		Password:    in.Password,
		Role:        in.Role,
	}

	if err := ur.checkUserExistsWithEmail(data.Email); err != nil {
		return err
	}

	if err := ur.checkUserExistsWithPhoneNumber(data.PhoneNumber); err != nil {
		return err
	}

	result := ur.db.GetDb().Create(data)

	if result.Error != nil {
		ur.logger.Error("error while inserting user", result.Error)
		return result.Error
	}

	ur.logger.Info(fmt.Sprintf("inserted user data:  %v", result.RowsAffected))

	return nil
}

func (ur *UserPostgresRepository) checkUserExistsWithEmail(email string) error {
	var user domain.User

	resultEmail := ur.db.GetDb().Where("email = ?", email).First(&user)
	if resultEmail.Error != nil {
		if errors.Is(resultEmail.Error, gorm.ErrRecordNotFound) {
			return nil
		}

		return resultEmail.Error
	}

	return error_custom.ErrUserAlreadyExistsWithEmail
}

func (ur *UserPostgresRepository) checkUserExistsWithPhoneNumber(phoneNumber string) error {
	var user domain.User

	resultEmail := ur.db.GetDb().Where("phone_number = ?", phoneNumber).First(&user)
	if resultEmail.Error != nil {
		if errors.Is(resultEmail.Error, gorm.ErrRecordNotFound) {
			return nil
		}

		return resultEmail.Error
	}

	return error_custom.ErrUserAlreadyExistsWithPhone
}

func (ur *UserPostgresRepository) UpdateUserData(in *domain.User) error {
	return nil
}

func (ur *UserPostgresRepository) DeleteUserData(uuid string) error {
	return nil
}

func (ur *UserPostgresRepository) FindUserData(filter map[string]interface{}) (*domain.User, error) {

	if len(filter) == 0 {
		return nil, fmt.Errorf("фильтр поиска не может быть пустым")
	}

	data := domain.User{}

	db := ur.db.GetDb().Model(&domain.User{})

	for key, value := range filter {
		db = db.Where(key+" = ?", value)
	}

	result := db.First(&data)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("пользователь не найден")
		}
		ur.logger.Error("ошибка при поиске пользователя", result.Error)
		return nil, result.Error
	}

	ur.logger.Info(fmt.Sprintf("найден пользователь: %v", filter))
	return &data, nil
}
