package server

import (
	"fmt"
	"github.com/Miroslovelife/whareflow/internal/config"
	http "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/handler"
	http2 "github.com/Miroslovelife/whareflow/internal/deliviry/http/v1/middleware"
	"github.com/Miroslovelife/whareflow/internal/repositories"
	"github.com/Miroslovelife/whareflow/internal/services"
	"github.com/Miroslovelife/whareflow/internal/usecase"
	"github.com/Miroslovelife/whareflow/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"log/slog"
)

type echoServer struct {
	app    *echo.Echo
	db     database.Database
	logger slog.Logger
	cfg    config.Config
}

func NewEchoServer(logger slog.Logger, db database.Database, cfg *config.Config) *echoServer {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		app:    echoApp,
		db:     db,
		logger: logger,
		cfg:    *cfg,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	// Health check adding
	s.app.GET("v1/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	s.initializeUserHttpHandler()

	serverUrl := fmt.Sprintf(":%d", 8089)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeUserHttpHandler() {

	userPostgresRepository := repositories.NewUserPostgresRepository(s.db, s.logger)

	passwordHahser := services.NewSHA1Hasher(s.cfg.Auth.PasswordSalt)

	manager := services.NewTokenM()

	userUsecase := usecase.NewUserUsecase(userPostgresRepository, passwordHahser, manager)

	userHttpHandler := http.NewUserHttpHandler(s.logger, userUsecase, s.cfg)

	authHttpMiddleware := http2.NewAuthHttpMiddleware(manager)

	// Routers
	userRouters := s.app.Group("v1")
	userRouters.POST("/reg", userHttpHandler.Register)
	userRouters.POST("/login_tel", userHttpHandler.LoginByPhoneNumber)
	userRouters.POST("/login_email", userHttpHandler.LoginByEmail)
	userRouters.Group("", authHttpMiddleware.Auth)

}
