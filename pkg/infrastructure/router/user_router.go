package router

import (
	"github.com/labstack/echo/v4"
	service "news/pkg/infrastructure/restful/service/user"
)

type UserRouter struct {
	serv service.UserService
}

func NewUserRouter(e *echo.Echo, userService *service.UserService) *echo.Echo {
	r := e.Group("/user")
	r.GET("/", userService.UserList)
	r.POST("/", userService.UserCreate)
	r.GET("/:id", userService.UserGetById)
	r.PUT("/:id", userService.UserUpdate)
	r.DELETE("/:id", userService.UserDelete)
	return e
}
