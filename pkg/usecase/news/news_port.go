package news

import "github.com/labstack/echo/v4"

type NewsInputPort interface {
	CreateNews(echo.Context, interface{}) (interface{}, error)
}

type NewsOutputPort interface {
	CreateNewsResponse(interface{}) (interface{}, error)
}
