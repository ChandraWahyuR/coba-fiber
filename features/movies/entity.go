package movies

import "github.com/gofiber/fiber/v2"

type MoviesHandlerInterface interface {
	// ViewAllMovies(c *fiber.Ctx) error
	RekomendasiMovies(c *fiber.Ctx) error
}
