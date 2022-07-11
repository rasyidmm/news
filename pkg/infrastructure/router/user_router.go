package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	service "news/pkg/infrastructure/restful/service/user"
)

type UserRouter struct {
	serv      service.UserService
	validator *validator.Validate
}

func (cv *UserRouter) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func NewUserRouter(e *echo.Echo, userService *service.UserService) *echo.Echo {
	e.Validator = &UserRouter{validator: validator.New()}
	r := e.Group("/user")
	r.GET("/", userService.UserList)
	r.POST("/", userService.UserCreate)
	r.GET("/:id", userService.UserGetById)
	r.PUT("/:id", userService.UserUpdate)
	r.DELETE("/:id", userService.UserDelete)
	return e
}
