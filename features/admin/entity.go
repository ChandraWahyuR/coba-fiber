package admin

import "github.com/gofiber/fiber/v2"

type Admin struct {
	ID          string
	Username    string
	Email       string
	Password    string
	FotoProfile string
	SuperAdmin  bool
}

type AdminHandlerInterface interface {
	GetAllUser(c *fiber.Ctx) error
}
type AdminDataInterface interface {
	GetAllUser(c *fiber.Ctx) error
}
type AdminServiceInterface interface {
	GetAllUser(c *fiber.Ctx) error
}
