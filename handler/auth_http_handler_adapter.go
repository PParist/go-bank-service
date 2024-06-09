package handler

import (
	"github.com/PParist/go-bank-service/entities"
	"github.com/PParist/go-bank-service/errorhandler"
	service "github.com/PParist/go-bank-service/services"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) authHandler {
	return authHandler{service: service}
}

func (h *authHandler) UserLogin(c *fiber.Ctx) error {
	userLogin := entities.UserLogin{}
	if err := c.BodyParser(&userLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := h.service.UserLogin(userLogin.Username, userLogin.Password)

	if err != nil {
		appErr, ok := err.(errorhandler.AppError)
		if ok {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": token})
}
