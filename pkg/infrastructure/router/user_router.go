package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	service "news/pkg/infrastructure/restful/service/user"
	"news/pkg/shared/jwtGen"
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
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/user")
	r.GET("/", userService.UserList, jwtGen.IsLoggedIn)
	r.POST("/", userService.UserCreate, jwtGen.IsLoggedIn)
	r.GET("/:id", userService.UserGetById, jwtGen.IsLoggedIn)
	r.PUT("/:id", userService.UserUpdate, jwtGen.IsLoggedIn)
	r.DELETE("/:id", userService.UserDelete, jwtGen.IsLoggedIn)
	return e
}
