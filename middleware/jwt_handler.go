package middleware

import (
	"presensi/helper"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/gofiber/jwt/v3"
)

func NewJWTMiddleware(secret string) fiber.Handler {
	return jwt.New(jwt.Config{
		SigningKey:   []byte(secret),
		ErrorHandler: helper.JWTErrorHandler,
	})
}
