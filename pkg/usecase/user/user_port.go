package user

import "github.com/labstack/echo/v4"

type UserInputPort interface {
	UserCreate(echo.Context, interface{}) (interface{}, error)
	UserList(echo.Context, interface{}) (interface{}, error)
	UserGetById(echo.Context, interface{}) (interface{}, error)
	UserUpdate(echo.Context, interface{}) (interface{}, error)
	UserDelete(echo.Context, interface{}) (interface{}, error)
}

type UserOutputPort interface {
	UserCreateResponse(interface{}) (interface{}, error)
	UserListResponse(interface{}) (interface{}, error)
	UserGetByIdResponse(interface{}) (interface{}, error)
	UserUpdateResponse(interface{}) (interface{}, error)
	UserDeleteResponse(interface{}) (interface{}, error)
}
