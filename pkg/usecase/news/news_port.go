package news

import "github.com/labstack/echo/v4"

type NewsInputPort interface {
	CreateNews(echo.Context, interface{}) (interface{}, error)
	GetAllNews(echo.Context, interface{}) (interface{}, error)
	GetByAllNews(echo.Context, interface{}) (interface{}, error)
	UpdateNews(echo.Context, interface{}) (interface{}, error)
	DeleteNews(echo.Context, interface{}) (interface{}, error)
}

type NewsOutputPort interface {
	CreateNewsResponse(interface{}) (interface{}, error)
	GetAllNewsResponse(interface{}) (interface{}, error)
	GetByAllNewResponse(interface{}) (interface{}, error)
	UpdateNewsResponse(interface{}) (interface{}, error)
	DeleteNewsResponse(interface{}) (interface{}, error)
}
