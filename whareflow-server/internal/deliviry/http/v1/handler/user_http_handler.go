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
	"time"
)

type UserHandler interface {
	Register(echo.Context) error
	LoginByPhoneNumber(echo.Context) error
	LoginByEmail(echo.Context) error
	Refresh(echo.Context) error
	GetProfile(echo.Context) error
	Logout(c echo.Context) error
}

type IUserHttpHandler struct {
	userUseCase usecase.UserUsecase
	logger      *slog.Logger
	cfg         config.Config
}

func NewIUserHttpHandler(logger slog.Logger, userUseCase usecase.UserUsecase, cfg config.Config) *IUserHttpHandler {
	return &IUserHttpHandler{
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
func (h *IUserHttpHandler) Register(c echo.Context) error {
	reqBody := new(delivery.UserReg)

	if err := c.Bind(reqBody); err != nil {
		h.logger.Error(fmt.Sprintf("Incorrect request body: %v", err))
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	fmt.Println(reqBody.Role)

	if reqBody.Role != "owner" && reqBody.Role != "employer" {
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
func (h *IUserHttpHandler) LoginByPhoneNumber(c echo.Context) error {
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

	refreshCookie := http.Cookie{
		Name:     "refresh-token",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.Auth.ExpRefreshToken * 60,
		Expires:  time.Now().Add(time.Duration(h.cfg.Auth.ExpRefreshToken) * 60),
	}

	c.SetCookie(&refreshCookie)

	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": accessToken,
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
func (h *IUserHttpHandler) LoginByEmail(c echo.Context) error {
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

	refreshCookie := http.Cookie{
		Name:     "refresh-token",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.Auth.ExpRefreshToken * 60,
		Expires:  time.Now().Add(time.Duration(h.cfg.Auth.ExpRefreshToken) * 60),
	}

	c.SetCookie(&refreshCookie)

	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": accessToken,
	})
}

// Refresh godoc
// @Summary Обновление пары токенов
// @Description Регистрирует нового пользователя
// @Tags auth
// @Accept			json
// @Produce		json
// @Failure 400 {object} map[string]string "error: invalid request body"
// @Failure 500 {object} map[string]string "error: internal server error"
// @Router /auth/refresh [get]
func (h *IUserHttpHandler) Refresh(c echo.Context) error {

	cookies := c.Cookies()
	for _, cookie := range cookies {
		fmt.Println("Cookie:", cookie.Name, "Value:", cookie.Value)
	}

	refreshToken, err := c.Cookie("refresh-token")
	if err != nil || refreshToken == nil {
		h.logger.Error("Refresh token cookie is missing or invalid")
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "refresh token is missing or invalid",
		})
	}

	newAccessToken, newRefreshToken, err := h.userUseCase.Refresh(refreshToken.Value, h.cfg.Auth.SecretAccessToken, h.cfg.Auth.SecretRefreshToken, h.cfg.Auth.ExpAccessToken, h.cfg.Auth.ExpRefreshToken)
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

	refreshCookie := http.Cookie{
		Name:     "refresh-token",
		Value:    newRefreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   h.cfg.Auth.ExpRefreshToken * 60,
		Expires:  time.Now().Add(time.Duration(h.cfg.Auth.ExpRefreshToken) * 60),
	}

	c.SetCookie(&refreshCookie)

	return c.JSON(http.StatusOK, map[string]string{
		"accessToken": newAccessToken,
	})
}

func (h *IUserHttpHandler) GetProfile(c echo.Context) error {
	userId := c.Get("x-user-id").(string)

	profile, err := h.userUseCase.GetProfile(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Can't get profile")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"profile": profile,
	})
}

func (h *IUserHttpHandler) Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "refresh-token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	})

	return c.JSON(http.StatusOK, map[string]string{"message": "Logout successful"})
}
