package routes

import (
	authHandler "go-carbon-tracker/handlers/auth"
	"go-carbon-tracker/middlewares"
	authRepo "go-carbon-tracker/repositories/auth"
	authUsecase "go-carbon-tracker/usecases/auth"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	loggerConfig := middlewares.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}

	loggerMiddleware := loggerConfig.Init()

	e.Use(loggerMiddleware)

	e.Use(middleware.Recover())

	rateLimiterConfig := middlewares.RateLimiterConfig{
		Rate:      10,
		Burst:     30,
		ExpiresIn: 3 * time.Minute,
	}

	rateLimiterMiddleware := rateLimiterConfig.Init()

	e.Use(rateLimiterMiddleware)

	customValidator := middlewares.InitValidator()

	e.Validator = customValidator

	e.Pre(middleware.RemoveTrailingSlash())

	jwtConfig := middlewares.JWTConfig{
		SecretKey:       os.Getenv("JWT_SCRET_KEY"),
		ExpiresDuration: 1,
	}

	// authMiddlewareConfig := jwtConfig.Init()

	authRoute(e, db, &jwtConfig)
}

func authRoute(e *echo.Echo, db *gorm.DB, jwtConfig *middlewares.JWTConfig) {
	authRepo := authRepo.NewAuthRepository(db)
	authUsecase := authUsecase.NewAuthUsecase(authRepo, jwtConfig)
	authHandler := authHandler.NewAuthHandler(authUsecase)

	wishlist := e.Group("/api/v1/auth")
	wishlist.POST("/login", authHandler.Login)
	wishlist.POST("/register", authHandler.Register)
}
