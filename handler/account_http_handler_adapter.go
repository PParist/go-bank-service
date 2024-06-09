package handler

import (
	"fmt"

	"github.com/PParist/go-bank-service/entities"
	"github.com/PParist/go-bank-service/errorhandler"
	"github.com/PParist/go-bank-service/logs"
	service "github.com/PParist/go-bank-service/services"
	"github.com/gofiber/fiber/v2"
)

type accountHandler struct {
	service service.AccountService
}

func NewAccountHandler(service service.AccountService) accountHandler {
	return accountHandler{service: service}
}

func (h *accountHandler) CreateAccount(c *fiber.Ctx) error {
	user_uid := c.Params("user_uid")
	accountBody := new(entities.NewAccountRequest)
	if err := c.BodyParser(accountBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	accountRespons, err := h.service.CreateAccount(user_uid, *accountBody)

	if err != nil {
		appErr, ok := err.(errorhandler.AppError)
		if ok {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": accountRespons})
}

func (h *accountHandler) GetAccounts(c *fiber.Ctx) error {
	fmt.Println("getAccounts")
	accounts, err := h.service.GetAccounts()
	if err != nil {
		appErr, ok := err.(errorhandler.AppError)
		if ok {
			logs.Error(appErr)
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": accounts})
}

func (h *accountHandler) GetAccountByUserUID(c *fiber.Ctx) error {
	user_uid := c.Params("user_uid")
	accounts, err := h.service.GetAccountByUserUID(user_uid)
	if err != nil {
		appErr, ok := err.(errorhandler.AppError)
		if ok {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": accounts})
}

func (h *accountHandler) UpdateAccount(c *fiber.Ctx) error {
	accountUid := c.Params("account_uid")
	accountBody := new(entities.AccountUpdateRequest)
	if err := c.BodyParser(accountBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	accounts, err := h.service.UpdateAccount(accountUid, *accountBody)
	if err != nil {
		appErr, ok := err.(errorhandler.AppError)
		if ok {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": accounts})
}

func (h *accountHandler) DeleteAccountByUID(c *fiber.Ctx) error {
	accountUid := c.Params("account_uid")
	err := h.service.DeleteAccountByUID(accountUid)
	if err != nil {
		appErr, ok := err.(errorhandler.AppError)
		if ok {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}
	}
	fmt.Println(err)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": "Delete sucress"})
}
