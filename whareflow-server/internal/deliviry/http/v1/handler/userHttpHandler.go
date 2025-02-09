package http

import (
	"errors"
	"fmt"
	"github.com/Miroslovelife/whareflow/internal/config"
	model_http "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	error_custom "github.com/Miroslovelife/whareflow/internal/errors"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type userHttpHandler struct {
	userUseCase usecase.UserUsecase
	logger      *slog.Logger
	cfg         config.Config
}

func NewUserHttpHandler(logger slog.Logger, userUseCase usecase.UserUsecase, cfg config.Config) *userHttpHandler {
	return &userHttpHandler{
		userUseCase: userUseCase,
		logger:      &logger,
		cfg:         cfg,
	}
}

func (h *userHttpHandler) Register(c echo.Context) error {
	reqBody := new(model_http.UserReg)

	if err := c.Bind(reqBody); err != nil {
		h.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	if err := h.userUseCase.Register(reqBody); err != nil {
		if errors.Is(err, error_custom.ErrUserAlreadyExistsWithEmail) {
			h.logger.Info("User already exists with email")
			return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
		}

		if errors.Is(err, error_custom.ErrUserAlreadyExistsWithPhone) {
			h.logger.Info("User already exists with phone number")
			return c.JSON(http.StatusOK, map[string]string{"error": err.Error()})
		}

		h.logger.Error("Unexpected error during registration: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "user registered successfully",
	})
}

func (h *userHttpHandler) LoginByPhoneNumber(c echo.Context) error {
	reqBody := new(model_http.UserLoginByPhoneNumber)

	if err := c.Bind(reqBody); err != nil {
		h.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	accessToken, refreshToken, err := h.userUseCase.LoginByPhoneNumber(reqBody, h.cfg.Auth.SecretAccessToken, h.cfg.Auth.SecretRefreshToken, h.cfg.Auth.ExpAccessToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access-token":  accessToken,
		"refresh-token": refreshToken,
	})
}

func (h *userHttpHandler) LoginByEmail(c echo.Context) error {
	reqBody := new(model_http.UserLoginByEmail)

	if err := c.Bind(reqBody); err != nil {
		h.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	accessToken, refreshToken, err := h.userUseCase.LoginByEmail(reqBody, h.cfg.Auth.SecretAccessToken, h.cfg.Auth.SecretRefreshToken, h.cfg.Auth.ExpAccessToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access-token":  accessToken,
		"refresh-token": refreshToken,
	})
}
