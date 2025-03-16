package repositories

import (
	"github.com/Miroslovelife/whareflow/internal/domain"
	"github.com/Miroslovelife/whareflow/pkg/database"
	"log/slog"
)

type PermissionRepository interface {
	CreatePermission(warehouseId uint, userId, roleName string, permissions []string) error
	GetPermission(warehouseId uint, userId, action string) (*domain.Permission, error)
	//GetAllPermissions()
	//GetRoles()
	//GetAllRoles()
}

type PermissionPostgresRepository struct {
	db     database.Database
	logger slog.Logger
}

func NewPermissionPostgresRepository(db database.Database, logger slog.Logger) *PermissionPostgresRepository {
	return &PermissionPostgresRepository{
		db:     db,
		logger: logger,
	}
}

func (pr *PermissionPostgresRepository) GetPermission(warehouseId uint, userId, action string) (*domain.Permission, error) {
	var permission domain.Permission

	err := pr.db.GetDb().
		Model(&domain.Permission{}).
		Joins("JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Joins("JOIN warehouse_user_roles ON warehouse_user_roles.role_id = role_permissions.role_id").
		Where("warehouse_user_roles.ware_house_id = ?", warehouseId).
		Where("warehouse_user_roles.user_id = ?", userId).
		Where("permissions.name = ?", action).
		First(&permission).Error

	if err != nil {
		return nil, err
	}

	return &permission, nil
}

func (pr *PermissionPostgresRepository) CreatePermission(warehouseId uint, userId, roleName string, permissions []string) error {
	tx := pr.db.GetDb().Begin()
	defer tx.Rollback()

	role := domain.Role{
		Name: roleName,
	}

	if err := tx.Create(role).Error; err != nil {
		return err
	}

	if err := tx.Create(&domain.WareHouseUserRole{
		WareHouseId: warehouseId,
		UserUuid:    userId,
		RoleId:      role.Id,
	}).Error; err != nil {
		return err
	}

	for _, perm_value := range permissions {
		var permission domain.Permission

		tx.Model(&domain.Permission{}).Where("name = ?", perm_value).First(&permission)

		if err := tx.Create(&domain.RolePermission{
			RoleId:       role.Id,
			PermissionId: permission.Id,
		}).Error; err != nil {
			return err
		}
	}

	return tx.Commit().Error
}
