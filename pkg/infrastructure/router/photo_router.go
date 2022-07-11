package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	service "news/pkg/infrastructure/restful/service/photo"
)

type PhotoRouter struct {
	photoService *service.PhotoService
	validator    *validator.Validate
}

func NewPhotoRouter(e *echo.Echo, photoService *service.PhotoService) *echo.Echo {
	e.Validator = &UserRouter{validator: validator.New()}
	r := e.Group("/photo")
	r.POST("/", photoService.PhotoCreate)
	r.GET("/", photoService.GetPhoto)
	r.PUT("/:id", photoService.UpdatePhoto)
	r.GET("/:id", photoService.PhotoGetById)
	r.DELETE("/:id", photoService.PhotoDeleteById)

	return e
}
