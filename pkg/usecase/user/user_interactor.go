package user

import (
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"news/pkg/domain/entity"
	"news/pkg/domain/repository"
	"news/pkg/shared/enum"
	"news/pkg/shared/tracing"
	"news/pkg/shared/util"
)

type UserCreateRequest struct {
	Username       string
	FirstName      string
	LastName       string
	Twitter        string
	Facebook       string
	Instagram      string
	Biography      string
	Email          string
	NomerHandphone string
	Password       string
	JenisUser      string
}
type UserCreateResponse struct {
	StatusCode string
	StatusDesc string
}
type UserListRequest struct {
	CurPage int
	Limit   int
}
type User struct {
	Id             string
	Username       string
	FirstName      string
	LastName       string
	Twitter        string
	Facebook       string
	Instagram      string
	Biography      string
	Email          string
	NomerHandphone string
	Password       string
	JenisUser      string
}

type UserListResponse struct {
	Data             []User
	PaginationHelper util.PaginationHelper
}

type UserGetByIdRequest struct {
	UserId string
}

type UserGetByIdResponse struct {
	Id             string
	Username       string
	FirstName      string
	LastName       string
	Twitter        string
	Facebook       string
	Instagram      string
	Biography      string
	Email          string
	NomerHandphone string
	Password       string
	JenisUser      string
	StatusCode     string
	StatusDesc     string
}
type UserUpdateRequest struct {
	Id             string
	Username       string
	FirstName      string
	LastName       string
	Twitter        string
	Facebook       string
	Instagram      string
	Biography      string
	Email          string
	NomerHandphone string
	Password       string
	JenisUser      string
}
type UserUpdateResponse struct {
	StatusCode string
	StatusDesc string
}

type UserDeleteRequest struct {
	UserId string
}

type UserDeleteResponse struct {
	StatusCode string
	StatusDesc string
}

type UserInteractor struct {
	repo repository.UserRepository
	out  UserOutputPort
}

func NewUserInteractor(r repository.UserRepository, o UserOutputPort) *UserInteractor {
	return &UserInteractor{
		repo: r,
		out:  o,
	}
}

func (i *UserInteractor) UserCreate(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}
	reqdata := in.(*UserCreateRequest)
	var request *entity.UserCreateRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}
	resData, err := i.repo.UserCreate(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var res *entity.UserCreateResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, http.ErrAbortHandler
	}

	tracing.LogResponse(sp, res)
	return i.out.UserCreateResponse(res)
}

func (i *UserInteractor) UserList(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}
	reqdata := in.(*UserListRequest)
	var request *entity.UserListRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}
	resData, err := i.repo.UserList(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var res *entity.UserListResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return i.out.UserListResponse(res)
}

func (i *UserInteractor) UserGetById(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}
	reqdata := in.(*UserGetByIdRequest)
	var request *entity.UserGetByIdRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}
	resData, err := i.repo.UserGetById(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var res *entity.UserGetByIdResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return i.out.UserGetByIdResponse(res)
}

func (i *UserInteractor) UserUpdate(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}
	reqdata := in.(*UserUpdateRequest)
	var request *entity.UserUpdateRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}
	resData, err := i.repo.UserUpdate(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var res *entity.UserUpdateResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return i.out.UserUpdateResponse(res)
}

func (i *UserInteractor) UserDelete(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}
	reqdata := in.(*UserDeleteRequest)
	var request *entity.UserDeleteRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}
	resData, err := i.repo.UserDelete(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var res *entity.UserDeleteResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, res)
	return i.out.UserDeleteResponse(res)
}
