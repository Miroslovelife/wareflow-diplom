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

type userPostgresRepository struct {
	db     database.Database
	logger slog.Logger
}

func NewUserPostgresRepository(db database.Database, logger slog.Logger) *userPostgresRepository {
	return &userPostgresRepository{
		db:     db,
		logger: logger,
	}
}

func (ur *userPostgresRepository) InsertUserData(in *domain.User) error {
	data := &domain.User{
		PhoneNumber: in.PhoneNumber,
		Username:    in.Username,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		Surname:     in.Surname,
		Email:       in.Email,
		Password:    in.Password,
		Role:        "user",
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

func (ur *userPostgresRepository) checkUserExistsWithEmail(email string) error {
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

func (ur *userPostgresRepository) checkUserExistsWithPhoneNumber(phoneNumber string) error {
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

func (ur *userPostgresRepository) UpdateUserData(in *domain.User) error {
	return nil
}

func (ur *userPostgresRepository) DeleteUserData(uuid string) error {
	return nil
}

func (ur *userPostgresRepository) FindUserData(filter map[string]interface{}) (*domain.User, error) {

	if len(filter) == 0 {
		return nil, fmt.Errorf("фильтр поиска не может быть пустым")
	}

	data := &domain.User{}

	db := ur.db.GetDb().Model(&domain.User{})

	for key, value := range filter {
		db = db.Where(key+" = ?", value)
	}

	result := db.First(data)
	fmt.Printf("%v", *data)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("пользователь не найден")
		}
		ur.logger.Error("ошибка при поиске пользователя", result.Error)
		return nil, result.Error
	}

	ur.logger.Info(fmt.Sprintf("найден пользователь: %v", filter))
	return data, nil
}

//func (ur *userPostgresRepository) CheckExistsUserData(filter map[string]interface{}) (bool, error) {
//	if len(filter) == 0 {
//		return false, fmt.Errorf("фильтр поиска не может быть пустым")
//	}
//
//	db := ur.db.GetDb()
//
//	for key, value := range filter {
//		db = db.Where(key+" = ?", value)
//	}
//
//	var count int64
//	result := db.Model(&domain.User{}).Limit(1).Count(&count)
//
//	if result.Error != nil {
//		ur.logger.Error("ошибка при проверке существования данных", result.Error)
//		return false, result.Error
//	}
//
//	return count > 0, nil
//}
