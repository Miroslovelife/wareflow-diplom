package http

import "github.com/labstack/echo/v4"

type UserHandler interface {
	Register(echo.Context) error
	LoginByPhoneNumber(echo.Context) error
	LoginByEmail(echo.Context) error
}
