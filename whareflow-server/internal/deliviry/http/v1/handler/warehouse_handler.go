package handler

import (
	"errors"
	"fmt"
	delivery "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	custom_errors "github.com/Miroslovelife/whareflow/internal/errors"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type WarehouseHandler interface {
	GetAllWarehouses(echo.Context) error
	GetWarehouse(echo.Context) error
	CreateWarehouse(echo.Context) error
	UpdateWarehouse(echo.Context) error
	DeleteWarehouse(c echo.Context) error
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
// @Router /warehouse [get]
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

	warehouseName := c.Param("name")

	warehouse, err := wh.whUsecase.GetWarehouse(userId, warehouseName)
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
		wh.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
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
	warehouseName := c.Param("name")
	userId := c.Get("x-user-id")
	if userId == nil {
		return c.JSON(http.StatusInternalServerError, "User ID Not Found")
	}
	userIdBytes := userId.(string)

	if err := wh.whUsecase.UpdateWarehouse(reqBody, warehouseName, userIdBytes); err != nil {
		wh.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		if errors.Is(err, custom_errors.ErrWarehouseAlreadyExist) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("You have already warehouse with name: %s", reqBody.Name),
			})
		}
		if errors.Is(err, custom_errors.ErrWareHouseNotFound) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("Warehouse not found with name: %s", warehouseName),
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
	warehouseName := c.Param("name")

	if err := wh.whUsecase.DeleteWarehouse(warehouseName, userId); err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"warehouses": fmt.Sprintf("warehouse success delete with name: %s", warehouseName),
	})

}
