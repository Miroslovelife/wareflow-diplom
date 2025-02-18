package custom_middleware

import (
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type roleHttpMiddleware struct {
	userUsecase usecase.UserUsecase
}

func NewRoleHttpMiddleware(userUsecase usecase.UserUsecase) *roleHttpMiddleware {
	return &roleHttpMiddleware{
		userUsecase: userUsecase,
	}
}

func (rm *roleHttpMiddleware) IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("x-user-id")

		ok, err := rm.userUsecase.IsAdmin(userId.(string))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "")
		}

		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "You do not have admin role")
		}

		return next(c)
	}
}

func (rm *roleHttpMiddleware) IsEmployer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Get("x-user-id")

		ok, err := rm.userUsecase.IsEmployer(userId.(string))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "")
		}

		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "You do not have admin role")
		}

		return next(c)
	}
}
