package http

import (
	"github.com/Miroslovelife/whareflow/internal/services"
	"github.com/labstack/echo/v4"
	"strings"
	"sync"
)

type authHttpMiddleware struct {
	tokenManager services.TokenManager
	mutex        sync.Mutex
}

func NewAuthHttpMiddleware(tokenManager services.TokenManager) *authHttpMiddleware {
	return &authHttpMiddleware{
		tokenManager: tokenManager,
	}
}

func (am *authHttpMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		am.mutex.Lock()
		defer am.mutex.Unlock()

		authHeader := c.Param("Authorization")

		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := am.tokenManager.IsAuthorized(authToken, "")
			if err != nil {
				c.Error(err)
			}
			if authorized {
				userId, err := am.tokenManager.ExtractUuidFromToken(authToken, "")
				if err != nil {
					c.Error(err)
				}

				c.Set("x-user-id", userId)
			}

		}
		return next(c)
	}
}
