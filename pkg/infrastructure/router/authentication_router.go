package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	service "news/pkg/infrastructure/restful/service/authentication"
)

type AuthenticationRouter struct {
	service.AuthenticationService
	validator *validator.Validate
}

func NewAuthenticationRouter(e *echo.Echo, authenticationService *service.AuthenticationService) *echo.Echo {
	e.Validator = &UserRouter{validator: validator.New()}
	r := e.Group("/authentication")
	r.POST("/login", authenticationService.Login)
	return e
}
