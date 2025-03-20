package handler

import (
	"errors"
	"fmt"
	delivery "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	custom_errors "github.com/Miroslovelife/whareflow/internal/errors"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"strconv"
)

type WarehouseHandler interface {
	GetAllWarehouses(echo.Context) error
	GetWarehouse(echo.Context) error
	CreateWarehouse(echo.Context) error
	UpdateWarehouse(echo.Context) error
	DeleteWarehouse(echo.Context) error
	GetEmployers(echo.Context) error
	GetWhsEmployer(c echo.Context) error
}

type IWareHouseHandler struct {
	logger    slog.Logger
	whUsecase usecase.WarehouseUsecase
}

func NewIWareHouseHandler(logger slog.Logger, whUsecase usecase.WarehouseUsecase) *IWareHouseHandler {
	return &IWareHouseHandler{
		logger:    logger,
		whUsecase: whUsecase,
	}
}

// GetAllWarehouses godoc
// @Summary Return warehouses list
// @Description Возвращает список всех складов пользователя
// @Tags warehouse
// @Accept			json
// @Produce		json
// @Success 200 {object} map[string]string "message: warehouse success created"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security ApiKeyAuth
// @Router /warehouse [get] admin/warehouse [get]
func (wh *IWareHouseHandler) GetAllWarehouses(c echo.Context) error {
	userId := c.Get("x-user-id").(string)

	warehouses, err := wh.whUsecase.GetAllWarehouse(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"warehouses": warehouses,
	})

}

// GetWarehouse godoc
// @Summary Return warehouses list
// @Description Возвращает список всех складов пользователя
// @Tags warehouse
// @Accept			json
// @Produce		json
// @Param			name	path		string	true	"warehouse name"
// @Success 200 {object} map[string]string "message: warehouse success created"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security ApiKeyAuth
// @Router /warehouse/{name} [get]
func (wh *IWareHouseHandler) GetWarehouse(c echo.Context) error {
	userId := c.Get("x-user-id").(string)

	warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	warehouse, err := wh.whUsecase.GetWarehouse(userId, uint(warehouseId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, warehouse)

}

// CreateWarehouse godoc
// @Summary Создание склада
// @Description Создает новый склад
// @Tags warehouse
// @Accept			json
// @Produce		json
// @Param request body delivery.WarehouseModelRequest true "Данные для создания склада"
// @Success 200 {object} map[string]string "message: warehouse success created"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security		ApiKeyAuth
// @Router /warehouse [post]
func (wh *IWareHouseHandler) CreateWarehouse(c echo.Context) error {
	reqBody := delivery.WarehouseModelRequest{}

	if err := c.Bind(&reqBody); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	userId := c.Get("x-user-id")
	if userId == nil {
		return c.JSON(http.StatusInternalServerError, "User ID Not Found")
	}
	userIdBytes := userId.(string)

	if err := wh.whUsecase.CreateWarehouse(reqBody, userIdBytes); err != nil {
		wh.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		if errors.Is(err, custom_errors.ErrWarehouseAlreadyExist) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("You have already warehouse with name: %s", reqBody.Name),
			})
		}
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "warehouse success created")
}

// UpdateWarehouse godoc
// @Summary Создание склада
// @Description Создает новый склад
// @Tags warehouse
// @Accept			json
// @Produce		json
// @Param			name	path		string	true	"warehouse name"
// @Param request body delivery.WarehouseModelRequest true "Данные для обновления склада"
// @Success 200 {object} map[string]string "message: warehouse success updated"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security		ApiKeyAuth
// @Router /warehouse/{name} [put]
func (wh *IWareHouseHandler) UpdateWarehouse(c echo.Context) error {
	reqBody := delivery.WarehouseModelRequest{}

	if err := c.Bind(&reqBody); err != nil {
		wh.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}
	warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}
	userId := c.Get("x-user-id")
	if userId == nil {
		return c.JSON(http.StatusInternalServerError, "User ID Not Found")
	}
	userIdBytes := userId.(string)

	if err := wh.whUsecase.UpdateWarehouse(reqBody, uint(warehouseId), userIdBytes); err != nil {
		wh.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		if errors.Is(err, custom_errors.ErrWarehouseAlreadyExist) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("You have already warehouse with name: %s", reqBody.Name),
			})
		}
		if errors.Is(err, custom_errors.ErrWareHouseNotFound) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("Warehouse not found with name: %s", warehouseId),
			})
		}
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "warehouse success created")
}

// DeleteWarehouse godoc
// @Summary Delete warehouse by name
// @Description Возвращает список всех складов пользователя
// @Tags warehouse
// @Accept			json
// @Produce		json
// @Param name path string true "warehouse name"
// @Success 200 {object} map[string]string "message: warehouse success delete"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security ApiKeyAuth
// @Router /warehouse/{name} [delete]
func (wh *IWareHouseHandler) DeleteWarehouse(c echo.Context) error {
	userId := c.Get("x-user-id").(string)
	warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	if err := wh.whUsecase.DeleteWarehouse(uint(warehouseId), userId); err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"warehouses": fmt.Sprintf("warehouse success delete with name: %s", warehouseId),
	})

}

// GetEmployers godoc
// @Summary Return warehouses employers
// @Description Возвращает список всех работников склада
// @Tags warehouse
// @Accept			json
// @Produce		json
// @Param			warehouse_id	path		int	true	"warehouse id"
// @Success 200 {object} map[string]string "[]delivery.Employer"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security ApiKeyAuth
// @Router /warehouse/{warehouse_id}/employer [get]
func (wh *IWareHouseHandler) GetEmployers(c echo.Context) error {
	warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't parse warehouse id")
	}

	userId := c.Get("x-user-id").(string)

	employers, err := wh.whUsecase.GetAllEmployers(uint(warehouseId), userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusForbidden, "You don't have permission on this warehouse")
		}
		return c.JSON(http.StatusInternalServerError, "Error when getting the list of employees")
	}

	return c.JSON(http.StatusOK, employers)
}

// GetWhsEmployer godoc
// @Summary Get a list of warehouses that can be accessed
// @Description Возвращает список всех складов до которых есть доступ у сотрудника
// @Tags warehouse
// @Accept			json
// @Produce		json
// @Success 200 {object} map[string]string "[]delivery.WarehouseModelResponse"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security ApiKeyAuth
// @Router /warehouse [get]
func (wh *IWareHouseHandler) GetWhsEmployer(c echo.Context) error {
	employerId := c.Get("x-user-id").(string)

	warehouses, err := wh.whUsecase.GetWhsEmployer(employerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Can't get warehouses")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
    		"warehouses": warehouses,
    	})
}
