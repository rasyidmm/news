package router

import (
	"github.com/labstack/echo/v4"
	service "news/pkg/infrastructure/restful/service/kategori"
)

// Authenticate godoc
// @Summary Authenticate
// @Description Authenticate
// @Tags auth
// @Accept json
// @Produce json
// @Param X-APP_KEY header string true "APP_KEY"
// @Param X-TRACE_ID header string true "uuid v4"
// @Param X-TIMESTAMP header string true "2006-01-02 15:04:05"
// @Param X-UUID header string true "UUID"
// @Param Payload body usecase.Authentication true "Authenticate"
// @Success 201 {object} utils.ResponseData{data=dto.AuthenticationResponseData}
// @Failure 400 {object} utils.ResponseData
// @Failure 401 {object} utils.ResponseData
// @Failure 422 {object} utils.ResponseData
// @Router /auth/external [post]
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
