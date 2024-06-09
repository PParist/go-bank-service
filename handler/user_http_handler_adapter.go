package handler

import (
	"fmt"
	"strconv"

	"github.com/PParist/go-bank-service/entities"
	"github.com/PParist/go-bank-service/errorhandler"
	service "github.com/PParist/go-bank-service/services"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) userHandler {
	return userHandler{service: service}
}

func (h *userHandler) CreateUsers(c *fiber.Ctx) error {
	user := entities.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.CreateUser(&user); err != nil {
		appErr, ok := err.(errorhandler.AppError)
		if ok {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})
}

func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	fmt.Println("get users")
	user, err := h.service.GetUsers()

	if err != nil {
		appErr, ok := err.(errorhandler.AppError)
		if ok {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})

}

func (h *userHandler) GetUserByID(c *fiber.Ctx) error {
	_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	user, err := h.service.GetUserByID(_id)

	if err != nil {
		appErr, ok := err.(errorhandler.AppError)
		if ok {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})

}

func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	user := entities.User{}
	_id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	//userToken := c.Locals(jwt.UserContextKey).(*entities.UserToken)
	if user, err := h.service.UpdateUserByID(_id, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": user})
	}

}
func (h *userHandler) DeleteUser(c *fiber.Ctx) error {
	_id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.service.DeleteUser(_id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	} else {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": "Delete sucress"})
	}

}
