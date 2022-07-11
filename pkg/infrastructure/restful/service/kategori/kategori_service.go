package kategori

import (
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"news/pkg/shared/tracing"
	usecase "news/pkg/usecase/kategori"
)

type KategoriCreateRequest struct {
	KategoriNama string `json:"kategorinama" validate:"required"`
	KategoriCode string `json:"kategoricode" validate:"required"`
}
type KategoriListRequest struct {
	CurPage int `json:"curPage"`
	Limit   int `json:"limit"`
}
type KategoriGetByIdRequest struct {
	KategoryId string `json:"KategoryId" validate:"required"`
}
type KategoriUpdateRequest struct {
	KategoriNama string `json:"kategorinama" validate:"required"`
	KategoriCode string `json:"kategoricode" validate:"required"`
}

type KategoriDeleteRequest struct {
	KategoryId string `json:"KategoryId" validate:"required"`
}

type KategoriService struct {
	uc usecase.KategoriInputPort
}

func NewKategoriService(u usecase.KategoriInputPort) *KategoriService {
	return &KategoriService{
		uc: u,
	}
}

func (s *KategoriService) KategoriCreate(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "KategoriCreate")
	defer sp.Finish()

	reqdata := new(KategoriCreateRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.KategoriCreateRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.KategoriCreate(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (s *KategoriService) KategoriList(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "KategoriList")
	defer sp.Finish()

	reqdata := new(KategoriListRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.KategoriListRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, c.JSON(http.StatusBadRequest, err.Error()))
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.KategoriList(c, request)
	if err != nil {
		tracing.LogError(sp, c.JSON(http.StatusBadRequest, err.Error()))
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (s *KategoriService) KategoriGetById(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "KategoriGetById")
	defer sp.Finish()

	id := c.Param("id")
	reqdata := &KategoriGetByIdRequest{
		KategoryId: id,
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.KategoriGetByIdRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.KategoriGetById(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (s *KategoriService) KategoriUpdate(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "KategoriUpdate")
	defer sp.Finish()

	id := c.Param("id")
	reqdata := new(KategoriUpdateRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	request := &usecase.KategoriUpdateRequest{
		Id:           id,
		KategoriNama: reqdata.KategoriNama,
		KategoriCode: reqdata.KategoriCode,
	}

	tracing.LogRequest(sp, request)

	res, err := s.uc.KategoriUpdate(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, err)
	return c.JSON(http.StatusOK, res)
}

func (s *KategoriService) KategoriDelete(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "KategoriDelete")
	defer sp.Finish()

	id := c.Param("id")
	reqdata := &KategoriDeleteRequest{
		KategoryId: id,
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.KategoriDeleteRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.KategoriDelete(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}
