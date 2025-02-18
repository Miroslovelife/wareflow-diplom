package custom_middleware

import "github.com/labstack/echo/v4"

type RoleMiddleware interface {
	IsAdmin(handlerFunc echo.HandlerFunc) echo.HandlerFunc
	IsEmployer(handlerFunc echo.HandlerFunc) echo.HandlerFunc
}
