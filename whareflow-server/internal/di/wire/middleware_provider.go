//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Miroslovelife/whareflow/internal/config"
	"github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/middleware"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/google/wire"
)

type MiddlewareProvider struct {
	AuthMiddleware *custom_middleware.AuthHttpMiddleware
	RoleMiddleware *custom_middleware.RoleHttpMiddleware
	WhMiddleware   *custom_middleware.IWhPermissionMiddleware
}

func AuthMiddlewareProvider(authUsecase usecase.AuthUsecase, cfg config.Config) *custom_middleware.AuthHttpMiddleware {
	return custom_middleware.NewAuthHttpMiddleware(authUsecase, cfg)
}

func RoleMiddlewareProvider(userUsecase usecase.UserUsecase) *custom_middleware.RoleHttpMiddleware {
	return custom_middleware.NewRoleHttpMiddleware(userUsecase)
}

func PermissionMiddlewareProvider(permissionUsecase usecase.PermissionUsecase) *custom_middleware.IWhPermissionMiddleware {
	return custom_middleware.NewWhPermissionMiddleware(permissionUsecase)
}

var MiddlewareProviderSet = wire.NewSet(
	AuthMiddlewareProvider,
	RoleMiddlewareProvider,
	PermissionMiddlewareProvider,
	wire.Struct(new(MiddlewareProvider), "AuthMiddleware", "RoleMiddleware", "WhMiddleware"),
)

func InitializeMiddlewareProviderSet(authUsecase usecase.AuthUsecase, cfg config.Config, userUsecase usecase.UserUsecase, permissionUsecase usecase.PermissionUsecase) MiddlewareProvider {
	wire.Build(MiddlewareProviderSet)
	return MiddlewareProvider{}
}
