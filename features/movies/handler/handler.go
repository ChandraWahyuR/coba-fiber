package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"presensi/constant"
	"presensi/features/movies"
	"presensi/helper"
	"presensi/utils"

	"github.com/gofiber/fiber/v2"
)

type HandlerMoves struct {
	j utils.JWTInterface
}

func NewMovieHandler(j utils.JWTInterface) movies.MoviesHandlerInterface {
	return &HandlerMoves{
		j: j,
	}
}

func (h *HandlerMoves) RekomendasiMovies(c *fiber.Ctx) error {
	tokenString := c.Get(constant.HeaderAuthorization)
	if tokenString == "" {
		return helper.UnauthorizedError(c)
	}
	ctx := c.Context()
	token, err := h.j.ValidateToken(ctx, tokenString)
	if err != nil {
		return helper.UnauthorizedError(c)
	}

	tokenData := h.j.ExtractUserToken(token)
	role, ok := tokenData[constant.JWT_ROLE]
	if !ok || (role != constant.RoleAdmin && role != constant.RoleUser) {
		return helper.UnauthorizedError(c)
	}
	movieTitle := c.Query("title")
	if movieTitle == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Judul film harus diberikan",
		})
	}

	pythonAPI := fmt.Sprintf("http://127.0.0.1:5000/recommend?title=%s", movieTitle)
	resp, err := http.Get(pythonAPI)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal menghubungi layanan rekomendasi",
		})
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal memproses respons API",
		})
	}

	return c.JSON(result)
}
