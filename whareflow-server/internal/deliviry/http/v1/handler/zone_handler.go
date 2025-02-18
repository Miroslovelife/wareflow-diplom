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

type ZoneHandler interface {
	CreateZone(echo.Context) error
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
// @Param request body delivery.ZoneModelRequest true "Данные для создания склада"
// @Success 200 {object} map[string]string "message: warehouse success created"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Security		ApiKeyAuth
// @Router /zone [post]
func (zh *IZoneHandler) CreateZone(c echo.Context) error {
	reqBody := delivery.ZoneModelRequest{}

	if err := c.Bind(&reqBody); err != nil {
		zh.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	if err := zh.zoneUsecase.CreateZone(reqBody); err != nil {
		zh.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		if errors.Is(err, custom_errors.ErrWarehouseAlreadyExist) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": fmt.Sprintf("You have already warehouse with name: %s", reqBody.Name),
			})
		}
		return c.JSON(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, "warehouse success created")
}
