package pkg

import (
	"github.com/PParist/go-bank-service/entities"
	"github.com/PParist/go-bank-service/errorhandler"
	service "github.com/PParist/go-bank-service/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const userContextKey = "user"

type jwtPackage struct {
	service service.JwtService
}

func NewJwtPackage(service service.JwtService) jwtPackage {
	return jwtPackage{service: service}
}

// Middleware to extract user data from JWT
func (p *jwtPackage) ValidateToken(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	if token == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userToken := new(entities.User)
	claims := token.Claims.(jwt.MapClaims)
	userToken.User_uid = claims["user_uid"].(string)
	userToken.User_Role = claims["role"].(string)

	// Store the user data in the Fiber context
	c.Locals(userContextKey, userToken)

	return c.Next()
}

// Middleware function to check user roles
func (p *jwtPackage) RoleRequired(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Missing token",
			})
		}
		user := c.Locals("user").(*entities.User)
		if err := p.service.ValidateRole(user.User_Role, roles); err != nil {
			appErr, ok := err.(errorhandler.AppError)
			if ok {
				return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
			}
		}
		return c.Next()
	}
}
