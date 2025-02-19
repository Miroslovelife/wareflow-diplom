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
	"strconv"
)

type ZoneHandler interface {
	CreateZone(echo.Context) error
	UpdateZone(c echo.Context) error
	GetAllZones(c echo.Context) error
	GetZone(c echo.Context) error
	DeleteZone(c echo.Context) error
}

type IZoneHandler struct {
	logger      slog.Logger
	zoneUsecase usecase.ZoneUsecase
}

func NewIZoneHandler(logger slog.Logger, zoneUsecase usecase.ZoneUsecase) *IZoneHandler {
	return &IZoneHandler{
		logger:      logger,
		zoneUsecase: zoneUsecase,
	}
}

// CreateZone godoc
// @Summary Создание зоны склада
// @Description Создает новую зону склада
// @Tags zone
// @Accept			json
// @Produce		json
// @Param warehouse_id	path		string	true	"warehouse id"
// @Param request body delivery.ZoneModelRequest true "Данные для создания склада"
// @Success 200 {object} map[string]string "message: warehouse success created"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security		ApiKeyAuth
// @Router /warehouse/{warehouse_id}/zone [post]
func (zh *IZoneHandler) CreateZone(c echo.Context) error {
	reqBody := delivery.ZoneModelRequest{}

	if err := c.Bind(&reqBody); err != nil {
		zh.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	userId := c.Get("x-user-id").(string)
	warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	if err := zh.zoneUsecase.CreateZone(reqBody, userId, warehouseId); err != nil {
		zh.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		if errors.Is(err, custom_errors.ErrWareHouseNotFound) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("warehouse not found: %s", reqBody.Name),
			})
		}
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "zone success created")
}

// UpdateZone godoc
// @Summary Обновление зоны склада
// @Description Обновляет зону склада
// @Tags zone
// @Accept			json
// @Produce		json
// @Param warehouse_id	path		string	true	"warehouse id"
// @Param zone_id	path		string	true	"zone id"
// @Param request body delivery.ZoneModelRequest true "Данные для обновления зоны склада"
// @Success 200 {object} map[string]string "message: zone success updated"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security		ApiKeyAuth
// @Router /warehouse/{warehouse_id}/zone/{zone_id} [put]
func (zh *IZoneHandler) UpdateZone(c echo.Context) error {
	reqBody := delivery.ZoneModelRequest{}

	if err := c.Bind(&reqBody); err != nil {
		zh.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	userId := c.Get("x-user-id").(string)

	zoneId, err := strconv.Atoi(c.Param("zone_id"))
	if err != nil {
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

	if err := zh.zoneUsecase.UpdateZone(reqBody, userId, zoneId, warehouseId); err != nil {
		zh.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		if errors.Is(err, custom_errors.ErrWareHouseNotFound) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("warehouse not found: %s", reqBody.Name),
			})
		}
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "zone success created")
}

// GetAllZones godoc
// @Summary Return zones list
// @Description Возвращает список всех зон склада
// @Tags zone
// @Accept			json
// @Produce		json
// @Param warehouse_id	path		string	true	"warehouse id"
// @Success 200 {object} map[string]string "message: warehouse success created"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security ApiKeyAuth
// @Router /warehouse/{warehouse_id}/zone [get]
func (zh *IZoneHandler) GetAllZones(c echo.Context) error {
	userId := c.Get("x-user-id").(string)

	warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	zones, err := zh.zoneUsecase.GetAllZone(userId, warehouseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"zones": zones,
	})
}

// GetZone godoc
// @Summary Return zones list
// @Description Возвращает зону склада
// @Tags zone
// @Accept			json
// @Produce		json
// @Param warehouse_id	path		string	true	"warehouse id"
// @Param zone_id	path		string	true	"zone id"
// @Success 200 {object} map[string]string "message: warehouse success created"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security ApiKeyAuth
// @Router /warehouse/{warehouse_id}/zone/{zone_id} [get]
func (zh *IZoneHandler) GetZone(c echo.Context) error {
	userId := c.Get("x-user-id").(string)

	warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	zoneId, err := strconv.Atoi(c.Param("zone_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	zones, err := zh.zoneUsecase.GetZone(userId, warehouseId, zoneId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"zone": zones,
	})
}

// DeleteZone godoc
// @Summary Return zones list
// @Description Удаляет зону склада
// @Tags zone
// @Accept			json
// @Produce		json
// @Param warehouse_id	path		string	true	"warehouse id"
// @Param zone_id	path		string	true	"zone id"
// @Success 200 {object} map[string]string "message: warehouse success created"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security ApiKeyAuth
// @Router /warehouse/{warehouse_id}/zone/{zone_id} [delete]
func (zh *IZoneHandler) DeleteZone(c echo.Context) error {
	userId := c.Get("x-user-id").(string)

	warehouseId, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	zoneId, err := strconv.Atoi(c.Param("zone_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	err = zh.zoneUsecase.DeleteZone(userId, warehouseId, zoneId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "zone success deleted")
}
