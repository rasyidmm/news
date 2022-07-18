package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	service "news/pkg/infrastructure/restful/service/kategori"
	"news/pkg/shared/jwtGen"
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
	serv      service.KategoriService
	validator *validator.Validate
}

func NewKategoriRouter(e *echo.Echo, kategoriService *service.KategoriService) *echo.Echo {
	e.Validator = &UserRouter{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/kategori")
	r.POST("/getall", kategoriService.KategoriList, jwtGen.IsLoggedIn)
	r.POST("", kategoriService.KategoriCreate, jwtGen.IsLoggedIn)
	r.GET("/:id", kategoriService.KategoriGetById, jwtGen.IsLoggedIn)
	r.POST("/update", kategoriService.KategoriUpdate, jwtGen.IsLoggedIn)
	r.POST("/delete", kategoriService.KategoriDelete, jwtGen.IsLoggedIn)
	return e
}
