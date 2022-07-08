package photo

import (
	"github.com/labstack/echo/v4"
)

type PhotoInputPort interface {
	PhotoCreate(echo.Context, interface{}) (interface{}, error)
	PhotoGet(echo.Context, interface{}) (interface{}, error)
	PhotoUpdate(echo.Context, interface{}) (interface{}, error)
	PhotoGetById(echo.Context, interface{}) (interface{}, error)
	PhotoDeleteById(echo.Context, interface{}) (interface{}, error)
}
type PhotoOutputPort interface {
	PhotoCreateResponse(interface{}) (interface{}, error)
	PhotoGetResponse(interface{}) (interface{}, error)
	PhotoUpdateResponse(interface{}) (interface{}, error)
	PhotoGetByIdResponse(interface{}) (interface{}, error)
	PhotoDeleteByIdResponse(interface{}) (interface{}, error)
}
