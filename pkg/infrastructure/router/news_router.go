package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	service "news/pkg/infrastructure/restful/service/news"
	"news/pkg/shared/jwtGen"
)

type NewsRouter struct {
	photoService *service.NewsService
	validator    *validator.Validate
}

func NewNewsRouter(e *echo.Echo, newsService *service.NewsService) *echo.Echo {
	e.Validator = &UserRouter{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/news")
	r.POST("", newsService.CreateNews, jwtGen.IsLoggedIn)
	r.POST("/getall", newsService.GetAllNews, jwtGen.IsLoggedIn)
	r.POST("/getbyall", newsService.GetByAllNew, jwtGen.IsLoggedIn)
	r.POST("/update", newsService.UpdateNews, jwtGen.IsLoggedIn)
	r.POST("/delete", newsService.DeleteNews, jwtGen.IsLoggedIn)

	return e
}
