package handler

import (
	"errors"
	"fmt"
	"github.com/Miroslovelife/whareflow/internal/config"
	delivery "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/model"
	error_custom "github.com/Miroslovelife/whareflow/internal/errors"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type UserHandler interface {
	Register(echo.Context) error
	LoginByPhoneNumber(echo.Context) error
	LoginByEmail(echo.Context) error
	Refresh(c echo.Context) error
}

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

// Register godoc
// @Summary Регистрация пользователя
// @Description Регистрирует нового пользователя
// @Tags auth
// @Accept			json
// @Produce		json
// @Param request body delivery.UserReg true "Данные для регистрации"
// @Success 200 {object} map[string]string "message: user registered successfully"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Router /auth/sign-up [post]
func (h *userHttpHandler) Register(c echo.Context) error {
	reqBody := new(delivery.UserReg)

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

// LoginByPhoneNumber godoc
// @Summary Авторизация по номеру телефона
// @Description Регистрирует нового пользователя
// @Tags auth
// @Accept			json
// @Produce		json
// @Param request body delivery.UserLoginByPhoneNumber true "Данные для авторизации"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Router /auth/sign-in-phone [post]
func (h *userHttpHandler) LoginByPhoneNumber(c echo.Context) error {
	reqBody := new(delivery.UserLoginByPhoneNumber)

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

// LoginByEmail godoc
// @Summary Авторизация по почте
// @Description Регистрирует нового пользователя
// @Tags auth
// @Accept			json
// @Produce		json
// @Param request body delivery.UserLoginByEmail true "Данные для авторизации"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Router /auth/sign-in-email [post]
func (h *userHttpHandler) LoginByEmail(c echo.Context) error {
	reqBody := new(delivery.UserLoginByEmail)

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

// Refresh godoc
// @Summary Обновление пары токенов
// @Description Регистрирует нового пользователя
// @Tags auth
// @Accept			json
// @Produce		json
// @Param request body delivery.UserRefreshTokens true "ww"
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Router /auth/refresh [post]
func (h *userHttpHandler) Refresh(c echo.Context) error {
	reqBody := new(delivery.UserRefreshTokens)

	if err := c.Bind(reqBody); err != nil {
		h.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	accessToken, refreshToken, err := h.userUseCase.Refresh(reqBody.RefreshToken, h.cfg.Auth.SecretAccessToken, h.cfg.Auth.SecretRefreshToken, h.cfg.Auth.ExpAccessToken, h.cfg.Auth.ExpRefreshToken)
	if err != nil {
		if errors.Is(err, error_custom.ErrTokenIsNotValid) {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "token is not valid",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "invalid request body",
		})

	}

	return c.JSON(http.StatusOK, map[string]string{
		"access-token":  accessToken,
		"refresh-token": refreshToken,
	})
}
