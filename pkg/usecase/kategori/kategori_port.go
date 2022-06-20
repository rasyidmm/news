package kategori

import (
	"github.com/labstack/echo/v4"
)

type KategoriInputPort interface {
	KategoriCreate(echo.Context, interface{}) (interface{}, error)
	KategoriList(echo.Context, interface{}) (interface{}, error)
	KategoriGetById(echo.Context, interface{}) (interface{}, error)
	KategoriUpdate(echo.Context, interface{}) (interface{}, error)
	KategoriDelete(echo.Context, interface{}) (interface{}, error)
}

type KategoriOutputPort interface {
	KategoriCreateResponse(interface{}) (interface{}, error)
	KategoriListResponse(interface{}) (interface{}, error)
	KategoriGetByIdResponse(interface{}) (interface{}, error)
	KategoriUpdateResponse(interface{}) (interface{}, error)
	KategoriDeleteResponse(interface{}) (interface{}, error)
}
