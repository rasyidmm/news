package kategori

import (
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"news/pkg/domain/entity"
	"news/pkg/domain/repository"
	"news/pkg/shared/enum"
	"news/pkg/shared/jwtGen"
	"news/pkg/shared/tracing"
	"news/pkg/shared/util"
)

type KategoriCreateRequest struct {
	KategoriNama string
	KategoriCode string
}
type KategoriCreateResponse struct {
	StatusCode string
	StatusDesc string
}
type KategoriListRequest struct {
	CurPage int
	Limit   int
}
type Kategory struct {
	Id           string
	KategoriNama string
	KategoriCode string
}

type KategoriListResponse struct {
	Data             []Kategory
	PaginationHelper util.PaginationHelper
}
type KategoriGetByIdRequest struct {
	KategoryId string
}
type KategoriGetByIdResponse struct {
	Id           string
	KategoriNama string
	KategoriCode string
	StatusCode   string
	StatusDesc   string
}
type KategoriUpdateRequest struct {
	Id           string
	KategoriNama string
	KategoriCode string
}
type KategoriUpdateResponse struct {
	StatusCode string
	StatusDesc string
}

type KategoriDeleteRequest struct {
	KategoryId string
}

type KategoriDeleteResponse struct {
	StatusCode string
	StatusDesc string
}

type KatergoriInteractor struct {
	repo repository.KategoriRepository
	out  KategoriOutputPort
}

func NewKatergoriInteractor(r repository.KategoriRepository, o KategoriOutputPort) *KatergoriInteractor {
	return &KatergoriInteractor{
		repo: r,
		out:  o,
	}
}

func (s *KatergoriInteractor) KategoriCreate(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		return nil, status.Error(http.StatusBadRequest, "request nil")
	}

	reqdata := in.(*KategoriCreateRequest)
	var request *entity.KategoriCreateRequest

	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, http.ErrBodyNotAllowed
	}

	resData, err := s.repo.KategoriCreate(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var res *KategoriCreateResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, http.ErrAbortHandler
	}

	tracing.LogResponse(sp, res)
	return s.out.KategoriCreateResponse(res)
}
func (s *KatergoriInteractor) KategoriList(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request nill"))
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}
	reqdata := in.(*KategoriListRequest)
	var request *entity.KategoriListRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}
	resGet := jwtGen.GetClientMetadata(ctx)
	tracing.LogObject(sp, "IsiJwt", resGet)

	resData, err := s.repo.KategoriList(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var res *KategoriListResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	tracing.LogResponse(sp, res)
	return s.out.KategoriListResponse(res)
}
func (s *KatergoriInteractor) KategoriGetById(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request nill"))
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}
	reqdata := in.(*KategoriGetByIdRequest)
	var request *entity.KategoriGetByIdRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}
	resData, err := s.repo.KategoriGetById(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	var res *KategoriGetByIdResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, http.ErrAbortHandler
	}

	tracing.LogResponse(sp, res)
	return s.out.KategoriGetByIdResponse(res)
}
func (s *KatergoriInteractor) KategoriUpdate(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request nill"))
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}
	reqdata := in.(*KategoriUpdateRequest)
	var request *entity.KategoriUpdateRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}
	resData, err := s.repo.KategoriUpdate(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	var res *KategoriUpdateResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return s.out.KategoriUpdateResponse(res)
}
func (s *KatergoriInteractor) KategoriDelete(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}
	reqdata := in.(*KategoriDeleteRequest)
	var request *entity.KategoriDeleteRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}
	resData, err := s.repo.KategoriDelete(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var res *entity.KategoriDeleteResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, http.ErrAbortHandler
	}

	tracing.LogResponse(sp, res)
	return s.out.KategoriDeleteResponse(res)
}
