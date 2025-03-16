package custom_middleware

import (
	delivery "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WhPermissionMiddleware interface {
	HasPermissionOnWarehouse(handlerFunc echo.HandlerFunc) echo.HandlerFunc
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
		warehouseId := c.Get("x-warehouse-id")
		action := c.Param("action")

		permissionData := delivery.Permission{
			Uuid:        userId.(string),
			WareHouseId: warehouseId.(uint),
			Action:      action,
		}

		ownerId, err := wp.permissionUsecase.CheckPermission(permissionData)
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Access denied"})
		}

		c.Set("x-user-id", ownerId)

		return next(c)
	}

}
