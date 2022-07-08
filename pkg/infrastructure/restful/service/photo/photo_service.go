package photo

import (
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"news/pkg/shared/tracing"
	usecase "news/pkg/usecase/photo"
)

type CreatePhotoRequest struct {
	FileBase64  string `json:"filebase64"`
	FileName    string `json:"filename"`
	TypeFile    string `json:"typefile"`
	Description string `json:"description"`
}

type PhotoListRequest struct {
	CurPage   int    `json:"curPage"`
	Limit     int    `json:"limit"`
	NamePhoto string `json:"namePhoto"`
}
type UpdatePhotoRequest struct {
	Id          string `json:"id"`
	FileBase64  string `json:"filebase64"`
	FileName    string `json:"filename"`
	TypeFile    string `json:"typefile"`
	Description string `json:"description"`
}

type DeletePhotoRequest struct {
	Id string `json:"id"`
}

type PhotoGetByIdRequest struct {
	Id string `json:"id"`
}

type PhotoService struct {
	uc usecase.PhotoInputPort
}

func NewPhotoService(u usecase.PhotoInputPort) *PhotoService {
	return &PhotoService{
		uc: u,
	}
}
func (s *PhotoService) PhotoCreate(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "PhotoCreate")
	defer sp.Finish()

	reqdata := new(CreatePhotoRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.CreatePhotoRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.PhotoCreate(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (s *PhotoService) GetPhoto(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "GetPhoto")
	defer sp.Finish()

	reqdata := new(PhotoListRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.PhotoListRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.PhotoGet(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (s *PhotoService) UpdatePhoto(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "UpdatePhoto")
	defer sp.Finish()

	id := c.Param("id")
	reqdata := new(UpdatePhotoRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	request := &usecase.UpdatePhotoRequest{
		Id:          id,
		FileBase64:  reqdata.FileBase64,
		FileName:    reqdata.FileName,
		TypeFile:    reqdata.TypeFile,
		Description: reqdata.Description,
	}

	tracing.LogRequest(sp, request)

	res, err := s.uc.PhotoUpdate(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (s *PhotoService) PhotoGetById(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "PhotoGetById")
	defer sp.Finish()

	id := c.Param("id")
	reqdata := &PhotoGetByIdRequest{
		Id: id,
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.PhotoGetByIdRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)
	res, err := s.uc.PhotoGetById(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (s *PhotoService) PhotoDeleteById(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "PhotoGetById")
	defer sp.Finish()

	id := c.Param("id")
	reqdata := &DeletePhotoRequest{
		Id: id,
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.DeletePhotoRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)
	res, err := s.uc.PhotoDeleteById(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}
