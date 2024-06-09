package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PParist/go-bank-service/entities"
	"github.com/PParist/go-bank-service/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//const userContextKey = "user"

// const (
// 	host     = "103.117.148.23" // or the Docker service name if running in another container
// 	port     = 5432             // default PostgreSQL port
// 	user     = "myuser"         // as defined in docker-compose.yml
// 	password = "mypassword"     // as defined in docker-compose.yml
// 	dbname   = "mydatabase"     // as defined in docker-compose.yml
// )

func main() {
	initConfig()

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	db := initDatabase()

	db.AutoMigrate(&entities.User{}, &entities.Profile{}, &entities.Account{}, &entities.Role{}, &entities.Permission{}, &entities.RolePermission{})

	fmt.Println("Connected DB!")

	app := fiber.New()

	router.InitRouterConfig(app, db)

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil { // อ่านไฟล์ config
		fmt.Printf("Error reading config file, %s", err)
	}
}

func initTimeZone() {
	if ict, err := time.LoadLocation("Asia/Bangkok"); err != nil {
		panic(err)
	} else {
		time.Local = ict
	}
}

func initDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		viper.GetString("db.host"),
		viper.GetInt64("db.port"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"))

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log ALL level
			Colorful:      true,        // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect to database")
	}
	postgres, err := db.DB()
	if err != nil {
		panic("failed to connect to database")
	}
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	postgres.SetMaxOpenConns(10)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	postgres.SetConnMaxLifetime(3 * time.Minute)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	postgres.SetMaxIdleConns(10)

	return db
}
