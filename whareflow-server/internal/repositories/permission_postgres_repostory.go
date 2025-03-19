package repositories

import (
	"errors"
	"fmt"
	"github.com/Miroslovelife/whareflow/internal/domain"
	"github.com/Miroslovelife/whareflow/pkg/database"
	"gorm.io/gorm"
	"log/slog"
)

type PermissionRepository interface {
	CreateRole(warehouseId uint, ownerId, employerName, roleName string, permissions []uint) error
	GetPermission(warehouseId uint, userId, action string) (*domain.Permission, error)
	GetAllPermissions() (*[]domain.Permission, error)
	GetAllEmployerPermissions(warehouseId uint, ownerId, employerId string) (*[]domain.Permission, error)
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
		Model(&domain.Permission{}).Debug().
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

func (pr *PermissionPostgresRepository) CreateRole(warehouseId uint, ownerId, employerName, roleName string, permissions []uint) error {
	tx := pr.db.GetDb().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Проверяем, есть ли склад и принадлежит ли он owner'у
	var warehouse domain.WareHouse
	if err := tx.Where("id = ? AND uuid_user = ?", warehouseId, ownerId).First(&warehouse).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Проверяем, существует ли пользователь
	var employer domain.User
	if err := tx.Where("username = ?", employerName).First(&employer).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Проверяем, существует ли уже такая роль
	var role domain.Role
	err := tx.Where("name = ?", roleName).First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Если роли нет — создаём
			role = domain.Role{Name: roleName}
			if err := tx.Create(&role).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			tx.Rollback()
			return err
		}
	}

	// Проверяем, существует ли уже связь между пользователем, складом и ролью
	var warehouseUserRole domain.WarehouseUserRole
	err = tx.Where("ware_house_id = ? AND user_id = ? AND role_id = ?", warehouseId, employer.Uuid, role.Id).
		First(&warehouseUserRole).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Если связи нет — создаём новую
			newWarehouseUserRole := domain.WarehouseUserRole{
				WareHouseId: warehouseId,
				UserUuid:    string(employer.Uuid),
				RoleId:      role.Id,
			}
			if err := tx.Create(&newWarehouseUserRole).Error; err != nil {
				tx.Rollback()
				return err
			}
		} else {
			tx.Rollback()
			return err
		}
	}

	// Добавляем права к роли, если они ещё не привязаны
	for _, permValue := range permissions {
		var permission domain.Permission
		if err := tx.Where("id = ?", permValue).First(&permission).Error; err != nil {
			tx.Rollback()
			return err // Ошибка, если права не существует
		}

		// Проверяем, есть ли уже связь между ролью и привилегией
		var existingRolePermission domain.RolePermission
		err = tx.Where("role_id = ? AND permission_id = ?", role.Id, permission.Id).
			First(&existingRolePermission).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Если связи нет — создаём новую
				newRolePermission := domain.RolePermission{
					RoleId:       role.Id,
					PermissionId: permission.Id,
				}
				if err := tx.Create(&newRolePermission).Error; err != nil {
					tx.Rollback()
					return err
				}
			} else {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

func (pr *PermissionPostgresRepository) GetAllPermissions() (*[]domain.Permission, error) {
	var permissions []domain.Permission

	err := pr.db.GetDb().Model(&domain.Permission{}).Find(&permissions).Error
	if err != nil {
		return nil, err
	}

	return &permissions, nil
}

func (pr *PermissionPostgresRepository) GetAllEmployerPermissions(warehouseId uint, ownerId, employerId string) (*[]domain.Permission, error) {

	var count int64
	err := pr.db.GetDb().Model(&domain.WareHouse{}).Where("id = ? AND uuid_user = ?", warehouseId, ownerId).Count(&count).Error
	if err != nil {
		return nil, err
	}

	if count != 1 {
		return nil, fmt.Errorf("you don't have access on this warehouse")
	}

	var permissions []domain.Permission

	err = pr.db.GetDb().Model(&domain.Permission{}).
		Joins("JOIN role_permissions rp ON permissions.id = rp.permission_id").
		Joins("JOIN warehouse_user_roles wur ON rp.role_id = wur.role_id").
		Where("wur.ware_house_id = ? AND wur.user_id = ?", warehouseId, employerId).
		Find(&permissions).Error
	if err != nil {
		return nil, err
	}

	return &permissions, nil
}
