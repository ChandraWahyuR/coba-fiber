package config

import (
	"database/sql"

	userRepo "presensi/features/users/data"
	userHandler "presensi/features/users/handler"
	userService "presensi/features/users/services"

	moviesHandler "presensi/features/movies/handler"

	routes "presensi/route"
	"presensi/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type BootstrapConfig struct {
	DB  *sql.DB
	App *fiber.App
	Log *logrus.Logger
	JWT utils.JWTInterface
}

func Bootstrap(config *BootstrapConfig) {
	// Initialize Repositories
	usersRepo := userRepo.NewUserDataRepository(config.DB)
	usersService := userService.NewUserService(usersRepo, config.JWT)
	usersHandler := userHandler.NewUserHandler(usersService, config.JWT)

	// Movies
	moviesHandler := moviesHandler.NewMovieHandler(config.JWT)

	// Setup Routes
	routeConfig := routes.RouteConfig{
		App:                      config.App,
		UsersHandler:             usersHandler,
		RekomendasiMoviesHandler: moviesHandler,
	}

	// Initialize Routes
	routeConfig.Setup()
}
