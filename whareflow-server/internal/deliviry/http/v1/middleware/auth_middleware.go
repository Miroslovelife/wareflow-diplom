package custom_middleware

import (
	"github.com/Miroslovelife/whareflow/internal/config"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type AuthMiddleware interface {
	Auth(handlerFunc echo.HandlerFunc) echo.HandlerFunc
}

type AuthHttpMiddleware struct {
	authUsecase usecase.AuthUsecase
	cfg         config.Config
}

func NewAuthHttpMiddleware(authUsecase usecase.AuthUsecase, cfg config.Config) *AuthHttpMiddleware {
	return &AuthHttpMiddleware{
		authUsecase: authUsecase,
		cfg:         cfg,
	}
}

func (am *AuthHttpMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")
		authString := strings.Split(authHeader, " ")

		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Empty auth header")
		}

		if authString[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unknown authorization type")
		}

		auth, userId, err := am.authUsecase.Auth(authString[1], am.cfg.Auth.SecretAccessToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		if !auth {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid auth token")
		}
		c.Set("x-user-id", userId)
		return next(c)
	}
}
