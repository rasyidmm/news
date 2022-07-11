package authentication

import "github.com/labstack/echo/v4"

type AuthenticationInputPort interface {
	Login(echo.Context, interface{}) (interface{}, error)
}

type AuthenticationOutputPort interface {
	LoginResponse(interface{}) (interface{}, error)
}
