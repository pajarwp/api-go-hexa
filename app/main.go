package main

import (
	api "api-go-hexa/api"
	"api-go-hexa/config"
	"api-go-hexa/utils/mysql"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	logger "github.com/sirupsen/logrus"

	userControllerv1 "api-go-hexa/api/v1/user"
	userBusiness "api-go-hexa/business/user"
	userRepo "api-go-hexa/modules/user"
)

func main() {
	//init config
	config := config.GetConfigs()

	// init database
	dbCon := mysql.InitDatabase(&config.MySQLConfig)

	//initiate item repository
	userRepo := userRepo.NewRepository(dbCon)

	//initiate item service
	userService := userBusiness.NewUserService(userRepo)

	//initiate item controller
	userControllerV1 := userControllerv1.NewController(userService)

	// Echo instance
	e := echo.New()

	api.RegisterPath(e, userControllerV1)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
	}))

	// Start server
	go func() {
		address := fmt.Sprintf("localhost:%s", config.Port)

		if err := e.Start(address); err != nil {
			logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		logger.Fatal(err)
	}
}
