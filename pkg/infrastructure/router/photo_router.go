package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	service "news/pkg/infrastructure/restful/service/photo"
	"news/pkg/shared/jwtGen"
)

type PhotoRouter struct {
	photoService *service.PhotoService
	validator    *validator.Validate
}

func NewPhotoRouter(e *echo.Echo, photoService *service.PhotoService) *echo.Echo {
	e.Validator = &UserRouter{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/photo")
	r.POST("/", photoService.PhotoCreate, jwtGen.IsLoggedIn)
	r.GET("/", photoService.GetPhoto, jwtGen.IsLoggedIn)
	r.PUT("/:id", photoService.UpdatePhoto, jwtGen.IsLoggedIn)
	r.GET("/:id", photoService.PhotoGetById, jwtGen.IsLoggedIn)
	r.DELETE("/:id", photoService.PhotoDeleteById, jwtGen.IsLoggedIn)

	return e
}
