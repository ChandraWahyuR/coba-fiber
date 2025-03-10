package handler

import (
	"presensi/constant"
	"presensi/features/users"
	"presensi/helper"
	"presensi/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserHandler struct {
	s users.UserServiceInterface
	j utils.JWTInterface
}

func NewUserHandler(u users.UserServiceInterface, j utils.JWTInterface) users.UserHandlerInterface {
	return &UserHandler{
		s: u,
		j: j,
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var reqRegister UserRegisterRequest
	if err := c.BodyParser(&reqRegister); err != nil {
		status, message := helper.HandleFiberError(c, err)
		return c.Status(status).JSON(helper.FormatResponse(false, message, nil))
	}

	user := users.User{
		ID:              uuid.New().String(),
		NIM:             reqRegister.NIM,
		Email:           reqRegister.Email,
		Password:        reqRegister.Password,
		ConfirmPassword: reqRegister.ConfirmPassword,
	}

	err := h.s.Register(user)
	if err != nil {
		return c.Status(helper.ConverResponse(err)).JSON(helper.FormatResponse(false, err.Error(), nil))
	}

	return c.Status(fiber.StatusCreated).JSON(helper.FormatResponse(true, constant.RegisterBerhasil, nil))
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var reqLogin UserLoginRequest
	if err := c.BodyParser(&reqLogin); err != nil {
		status, message := helper.HandleFiberError(c, err)
		return c.Status(status).JSON(helper.FormatResponse(false, message, nil))
	}

	user := users.User{
		Email:    reqLogin.Email,
		Password: reqLogin.Password,
	}

	userData, err := h.s.Login(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.FormatResponse(false, err.Error(), nil))
	}

	response := UserLoginResponse{
		Token: userData.Token,
	}

	return c.Status(fiber.StatusOK).JSON(helper.FormatResponse(true, constant.LoginBerhasil, response))
}
