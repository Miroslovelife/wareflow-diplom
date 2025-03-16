//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/Miroslovelife/whareflow/internal/services"
	"github.com/Miroslovelife/whareflow/pkg/qr"
	"github.com/google/wire"
	"log/slog"
)

type ProviderService struct {
	TokenManager *services.TokenM
	Hasher       *services.SHA1Hasher
	QR           *qr.Generator
}

func ProvideTokenManagerService() *services.TokenM {
	return services.NewTokenM()
}

func ProvideHasherService(salt string) *services.SHA1Hasher {
	return services.NewSHA1Hasher(salt)
}

func ProvideQRService(logger slog.Logger) *qr.Generator {
	return qr.NewGenerator(logger)
}

var ServiceProviderSet = wire.NewSet(
	ProvideTokenManagerService,
	ProvideHasherService,
	ProvideQRService,
	wire.Struct(new(ProviderService), "TokenManager", "Hasher", "QR"),
)

func InitializeServiceProviderSet(salt string, logger slog.Logger) ProviderService {
	wire.Build(ServiceProviderSet)
	return ProviderService{}
}
