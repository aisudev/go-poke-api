package main

import (
	"fmt"
	"net/http"
	"poke/domain"
	"poke/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// environment
	utils.ViperInit()

	// connection to db
	DBConnection()

	// Auto Migration
	AutoMigration()
}

func main() {
	e := echo.New()

	// Middlewares
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server is running")
	})

	e.Logger.Fatal(
		e.Start(
			fmt.Sprintf(":%s", viper.GetString("app.port")),
		),
	)
}

func DBConnection() {
	var err error

	DB, err = gorm.Open(sqlite.Open("db/poke.db"), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("fatal error db connection: %s \n", err))
	}

	fmt.Println("DB is established...")
}

func AutoMigration() {
	DB.AutoMigrate(&domain.User{}, &domain.Poke{})
}
