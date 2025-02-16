package repositories

import (
	"github.com/Miroslovelife/whareflow/internal/domain"
)

type UserRepository interface {
	InsertUserData(in *domain.User) error
	UpdateUserData(in *domain.User) error
	FindUserData(filter map[string]interface{}) (*domain.User, error)
	DeleteUserData(uuid string) error
}
