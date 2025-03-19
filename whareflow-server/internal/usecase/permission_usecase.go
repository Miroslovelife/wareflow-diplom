package usecase

import (
	delivery "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	"github.com/Miroslovelife/whareflow/internal/repositories"
)

type PermissionUsecase interface {
	CheckPermission(in delivery.Permission) (string, error)
	CreateRole(in *delivery.RoleReq, warehouseId uint, ownerId string) error
	GetAllPermissions() (*[]delivery.PermissionResponse, error)
	GetAllUserPermission(warehouseId uint, ownerId, employerName string) (*[]delivery.PermissionResponse, error)
}

type IPermissionUsecase struct {
	permissionRepo repositories.PermissionRepository
	wareHouseRepo  repositories.WareHouseRepository
	userRepo       repositories.UserRepository
}

func NewIPermissionUsecase(userRepo repositories.UserRepository, permissionRepo repositories.PermissionRepository, wareHouseRepo repositories.WareHouseRepository) *IPermissionUsecase {
	return &IPermissionUsecase{
		userRepo:       userRepo,
		permissionRepo: permissionRepo,
		wareHouseRepo:  wareHouseRepo,
	}
}

func (pu *IPermissionUsecase) CheckPermission(in delivery.Permission) (string, error) {

	_, err := pu.permissionRepo.GetPermission(in.WareHouseId, in.Uuid, in.Action)
	if err != nil {
		return "", err
	}

	ownerId, err := pu.wareHouseRepo.FindWareHouseOwner(in.WareHouseId)
	if err != nil {
		return "", err
	}

	return ownerId.UuidUser, err
}

func (pu *IPermissionUsecase) CreateRole(in *delivery.RoleReq, warehouseId uint, ownerId string) error {
	err := pu.permissionRepo.CreateRole(warehouseId, ownerId, in.UserName, in.Name, in.Permissions)
	if err != nil {
		return err
	}

	return nil
}

func (pu *IPermissionUsecase) GetAllPermissions() (*[]delivery.PermissionResponse, error) {
	permissions, err := pu.permissionRepo.GetAllPermissions()
	if err != nil {
		return nil, err
	}

	var permissionsRes []delivery.PermissionResponse
	for _, permission := range *permissions {
		permissionRes := delivery.PermissionResponse{
			Id:   permission.Id,
			Name: permission.Name,
		}

		permissionsRes = append(permissionsRes, permissionRes)
	}

	return &permissionsRes, nil
}

func (pu *IPermissionUsecase) GetAllUserPermission(warehouseId uint, ownerId, employerName string) (*[]delivery.PermissionResponse, error) {
	findData := map[string]interface{}{
		"username": employerName,
	}

	employer, err := pu.userRepo.FindUserData(findData)
	if err != nil {
		return nil, err
	}

	permissions, err := pu.permissionRepo.GetAllEmployerPermissions(warehouseId, ownerId, string(employer.Uuid))
	if err != nil {
		return nil, err
	}

	var permissionsRes []delivery.PermissionResponse
	for _, permission := range *permissions {
		permissionRes := delivery.PermissionResponse{
			Id:   permission.Id,
			Name: permission.Name,
		}

		permissionsRes = append(permissionsRes, permissionRes)
	}

	return &permissionsRes, nil
}
