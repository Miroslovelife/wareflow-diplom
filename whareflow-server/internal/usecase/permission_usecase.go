package usecase

import (
	delivery "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	"github.com/Miroslovelife/whareflow/internal/repositories"
)

type PermissionUsecase interface {
	CheckPermission(in delivery.Permission) (string, error)
	//CreatePermission(in *delivery.Permission) error
}

type IPermissionUsecase struct {
	permissionRepo repositories.PermissionRepository
	wareHouseRepo  repositories.WareHouseRepository
}

func NewIPermissionUsecase(permissionRepo repositories.PermissionRepository, wareHouseRepo repositories.WareHouseRepository) *IPermissionUsecase {
	return &IPermissionUsecase{
		permissionRepo: permissionRepo,
		wareHouseRepo:  wareHouseRepo,
	}
}

func (pu *IPermissionUsecase) CheckPermission(in delivery.Permission) (string, error) {

	action, err := pu.permissionRepo.GetPermission(in.WareHouseId, in.Uuid, in.Action)
	if err != nil {
		return "", err
	}
	if action.Name != in.Action {
		return "", err
	}

	ownerId, err := pu.wareHouseRepo.FindWareHouseOwner(in.WareHouseId)
	if err != nil {
		return "", err
	}

	return ownerId.UuidUser, err
}
