package routes

import (
	"presensi/features/movies"
	"presensi/features/users"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                      *fiber.App
	UsersHandler             users.UserHandlerInterface
	RekomendasiMoviesHandler movies.MoviesHandlerInterface
	// AuthMiddleware fiber.Handler // karena validasinya di handler dan service
}

func (c *RouteConfig) Setup() {
	c.SetupUserRoute()
	c.SetupRekomendasiMovieRoute()
}
func (c *RouteConfig) SetupUserRoute() {
	// Auth di fiber
	// jwtMiddleware := middleware.NewJWTMiddleware("a123")
	// c.App.Use(jwtMiddleware)

	c.App.Post("/api/register", c.UsersHandler.Register)
	c.App.Post("/api/login", c.UsersHandler.Login)
}

func (c *RouteConfig) SetupRekomendasiMovieRoute() {
	// Auth di fiber
	// jwtMiddleware := middleware.NewJWTMiddleware("a123")
	// c.App.Use(jwtMiddleware)

	c.App.Get("/api/reckomendasi-movies", c.RekomendasiMoviesHandler.RekomendasiMovies)
}
