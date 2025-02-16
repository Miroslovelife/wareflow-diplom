package custom_middleware

import "github.com/labstack/echo/v4"

type AuthMiddleware interface {
	Auth(handlerFunc echo.HandlerFunc) echo.HandlerFunc
}
