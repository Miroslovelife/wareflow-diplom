package custom_middleware

import (
	delivery "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"fmt"
)

type WhPermissionMiddleware interface {
	HasPermissionOnWarehouse(handlerFunc echo.HandlerFunc) echo.HandlerFunc
	SetGroup(group string) echo.MiddlewareFunc
}

type IWhPermissionMiddleware struct {
	permissionUsecase usecase.PermissionUsecase
}

func NewWhPermissionMiddleware(permissionUsecase usecase.PermissionUsecase) *IWhPermissionMiddleware {
	return &IWhPermissionMiddleware{
		permissionUsecase: permissionUsecase,
	}
}

func (wp *IWhPermissionMiddleware) HasPermissionOnWarehouse(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("x-user-id")
		warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Invalid warehouse ID"})
		}

		// Получаем параметр action из пути
		action := c.Param("action")

		// Извлекаем группу из контекста
		group := c.Get("group").(string)


        fmt.Println(action)
		// Проверяем, что action соответствует нужному действию для этой группы
		if !wp.isValidActionForGroup(group, action) {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Invalid action for this group"})
		}

		// Формируем структуру данных для проверки прав доступа
		permissionData := delivery.Permission{
			Uuid:        userId.(string),
			WareHouseId: uint(warehouseId),
			Action:      action,
		}

		// Проверка прав доступа
		ownerId, err := wp.permissionUsecase.CheckPermission(permissionData)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Access denied"})
		}

		c.Set("x-user-id", ownerId)

		return next(c)
	}
}

func (wp *IWhPermissionMiddleware) isValidActionForGroup(group, action string) bool {
	switch group {
	case "self_perm":
        		if action != "get_my_permissions" {
        			return false
        		}
	case "warehouse":
    		if action != "warehouse_manage" {
    			return false
    		}
	case "zone":
		if action != "zone_manage" {
			return false
		}
	case "product":
		if action != "product_manage" {
			return false
		}
	case "role":
		if action != "role_manage" {
			return false
		}
	default:
		return false
	}
	return true
}

func (wp *IWhPermissionMiddleware) SetGroup(group string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("group", group)
			return next(c)
		}
	}
}
