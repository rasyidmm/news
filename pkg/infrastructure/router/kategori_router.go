package router

import (
	"github.com/labstack/echo/v4"
	service "news/pkg/infrastructure/restful/service/kategori"
)

type KategoriRouter struct {
	serv service.KategoriService
}

func NewKategoriRouter(e *echo.Echo, kategoriService *service.KategoriService) *echo.Echo {

	r := e.Group("/kategori")
	r.GET("/", kategoriService.KategoriList)
	r.POST("/", kategoriService.KategoriCreate)
	r.GET("/:id", kategoriService.KategoriGetById)
	r.PUT("/:id", kategoriService.KategoriUpdate)
	r.DELETE("/:id", kategoriService.KategoriDelete)
	return e
}
