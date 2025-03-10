package users

import "github.com/gofiber/fiber/v2"

type User struct {
	ID              string
	NIM             string
	Email           string
	Password        string
	ConfirmPassword string
}

type Login struct {
	ID       string
	Email    string
	Password string
	Token    string
}

type UserHandlerInterface interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type UserServiceInterface interface {
	Register(user User) error
	Login(user User) (Login, error)
}

type UserDataInterface interface {
	Register(user User) error
	Login(user User) (*Login, error)
}
