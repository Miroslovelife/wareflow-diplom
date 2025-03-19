package custom_middleware

import (
	"fmt"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RoleMiddleware interface {
	IsAdmin(next echo.HandlerFunc) echo.HandlerFunc
	IsOwner(handlerFunc echo.HandlerFunc) echo.HandlerFunc
	IsEmployer(handlerFunc echo.HandlerFunc) echo.HandlerFunc
}

type RoleHttpMiddleware struct {
	userUsecase usecase.UserUsecase
}

func NewRoleHttpMiddleware(userUsecase usecase.UserUsecase) *RoleHttpMiddleware {
	return &RoleHttpMiddleware{
		userUsecase: userUsecase,
	}
}

func (rm *RoleHttpMiddleware) IsOwner(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("x-user-id")
		fmt.Println(userId)
		ok, err := rm.userUsecase.IsOwner(userId.(string))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "")
		}

		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "You do not have owner role")
		}

		return next(c)
	}
}

func (rm *RoleHttpMiddleware) IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("x-user-id")

		ok, err := rm.userUsecase.IsAdmin(userId.(string))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "")
		}

		if !ok {
			return echo.NewHTTPError(http.StatusForbidden, "You do not have admin role")
		}

		return next(c)
	}
}

func (rm *RoleHttpMiddleware) IsEmployer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("x-user-id")

		ok, err := rm.userUsecase.IsEmployer(userId.(string))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "")
		}

		if !ok {
			return echo.NewHTTPError(http.StatusForbidden, "You do not have employer role")
		}

		return next(c)
	}
}
