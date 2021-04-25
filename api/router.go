package http

import (
	"api-go-hexa/api/v1/user"

	"github.com/labstack/echo/v4"
)

//RegisterPath Registera V1 API path
func RegisterPath(e *echo.Echo, userController *user.Controller) {
	if userController == nil {
		panic("item controller cannot be nil")
	}

	//item
	userV1 := e.Group("v1/user")
	userV1.GET("/:id", userController.GetByID)
	userV1.POST("/register", userController.UserRegister)
	userV1.POST("/login", userController.UserLogin)
	userV1.PUT("/:id", userController.Update)
	userV1.DELETE("/:id", userController.Delete)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
