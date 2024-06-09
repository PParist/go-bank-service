package router

import (
	"fmt"
	"os"

	"github.com/PParist/go-bank-service/handler"
	"github.com/PParist/go-bank-service/logs"
	"github.com/PParist/go-bank-service/middleware"
	"github.com/PParist/go-bank-service/pkg"
	"github.com/PParist/go-bank-service/repositories"
	service "github.com/PParist/go-bank-service/services"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func InitRouterConfig(app *fiber.App, db *gorm.DB) {
	userRepositories := repositories.NewGormUserRepository(db)
	authRepositories := repositories.NewAuthRepository(db)
	accountRepositories := repositories.NewAccountRepositoryDB(db)
	jwtRepositories := repositories.NewJWTRepository(db)

	userService := service.NewUserService(userRepositories)
	authService := service.NewAuthService(authRepositories)
	accountService := service.NewAccountService(accountRepositories)
	jwtService := service.NewJWTService(jwtRepositories)

	http := handler.NewUserHandler(userService)
	authHttp := handler.NewAuthHandler(authService)
	accountHttp := handler.NewAccountHandler(accountService)

	jwtPackage := pkg.NewJwtPackage(jwtService)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Adjust this to be more restrictive if needed
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Post("/api/login", authHttp.UserLogin)

	// JWT Middleware & Middleware check role(restricted)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRECT"))},
	}), jwtPackage.ValidateToken)

	api := app.Group("/api", middleware.ApiMiddleware, middleware.RateLimiter(), middleware.Compression())

	admin := api.Group("/admin", jwtPackage.RoleRequired("Admin"))
	admin.Post("/user", http.CreateUsers)
	admin.Delete("/remove/:id", http.DeleteUser)
	admin.Post("/account/:user_uid", accountHttp.CreateAccount)
	admin.Delete("/account/:account_uid", accountHttp.DeleteAccountByUID)

	// admin.Put("/user/:id", http.UpdateUser)
	// admin.Put("/accounts/:account_uid", accountHttp.UpdateAccount)
	// admin.Get("/users", http.GetUsers)
	// admin.Get("/user/:id", http.GetUserByID)
	// admin.Get("/accounts", accountHttp.GetAccounts)
	// admin.Get("/accounts/:user_uid", accountHttp.GetAccountByUserUID)

	manager := api.Group("/manager", jwtPackage.RoleRequired("Admin", "Manager"))
	manager.Put("/user/:id", http.UpdateUser)
	manager.Put("/accounts/:account_uid", accountHttp.UpdateAccount)
	// manager.Get("/users", http.GetUsers)
	// manager.Get("/user/:id", http.GetUserByID)
	// manager.Get("/accounts", accountHttp.GetAccounts)
	// manager.Get("/accounts/:user_uid", accountHttp.GetAccountByUserUID)

	api.Get("/users", http.GetUsers)
	api.Get("/user/:id", http.GetUserByID)
	api.Get("/accounts", accountHttp.GetAccounts)
	api.Get("/accounts/:user_uid", accountHttp.GetAccountByUserUID)

	logs.Info("Service start on port :" + viper.GetString("app.port"))
	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
}
