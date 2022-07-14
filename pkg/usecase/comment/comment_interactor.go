package comment

import (
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"news/pkg/domain/entity"
	"news/pkg/domain/repository"
	"news/pkg/shared/enum"
	"news/pkg/shared/jwtGen"
	"news/pkg/shared/tracing"
	"news/pkg/shared/util"
)

type CreateCommentRequest struct {
	NewsId      string
	CommentText string
}

type CreateCommentResponse struct {
	StatusCode string
	StatusDesc string
}

type Comment struct {
	Id          string
	NewsId      string
	Username    string
	CommentText string
}

type GetAllCommentRequest struct {
	CurPage int
	Limit   int
}

type GetAllCommentResponse struct {
	Data             []Comment
	PaginationHelper util.PaginationHelper
}

type GetByAllCommentRequest struct {
	CurPage     int
	Limit       int
	NewsId      string
	Username    string
	CommentText string
}

type GetByAllCommentResponse struct {
	Data             []Comment
	PaginationHelper util.PaginationHelper
}

type UpdateCommentRequest struct {
	Id          string
	Username    string
	CommentText string
}
type UpdateCommentResponse struct {
	StatusCode string
	StatusDesc string
}

type DeleteCommentRequest struct {
	Id string
}

type DeleteCommentResponse struct {
	StatusCode string
	StatusDesc string
}

type CommentInteractor struct {
	repo repository.CommentRepository
	out  CommentOutputPort
}

func NewCommentInteractor(r repository.CommentRepository, o CommentOutputPort) *CommentInteractor {
	return &CommentInteractor{
		repo: r,
		out:  o,
	}
}

func (i *CommentInteractor) CreateComment(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request nill"))
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}
	reqdata := in.(*CreateCommentRequest)
	var request *entity.CreateCommentRequest

	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	resGet, errGer := jwtGen.GetClientMetadata(ctx)
	if errGer != nil {
		tracing.LogError(sp, errGer)
		return nil, errGer
	}

	tracing.LogObject(sp, "GetClientMetadata", resGet)

	request.Username = resGet.Username
	resData, err := i.repo.CreateComment(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var res *CreateCommentResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	tracing.LogResponse(sp, res)
	return i.out.CreateCommentResponse(res)
}

func (i *CommentInteractor) GetAllComment(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request nill"))
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}
	reqdata := in.(*GetAllCommentRequest)
	var request *entity.GetAllCommentRequest

	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}
	resGet, errGer := jwtGen.GetClientMetadata(ctx)
	if errGer != nil {
		tracing.LogError(sp, errGer)
		return nil, errGer
	}

	tracing.LogObject(sp, "GetClientMetadata", resGet)

	resData, err := i.repo.GetAllComment(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var res *GetAllCommentResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	tracing.LogResponse(sp, res)
	return i.out.GetAllCommentResponse(res)
}
func (i *CommentInteractor) GetByAllComment(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request nill"))
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}
	reqdata := in.(*GetByAllCommentRequest)
	var request *entity.GetByAllCommentRequest

	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	resGet, errGer := jwtGen.GetClientMetadata(ctx)
	if errGer != nil {
		tracing.LogError(sp, errGer)
		return nil, errGer
	}

	tracing.LogObject(sp, "GetClientMetadata", resGet)

	request.Username = resGet.Username
	resData, err := i.repo.GetByAllComment(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var res *GetByAllCommentResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	tracing.LogResponse(sp, res)
	return i.out.GetByAllCommentResponse(res)
}
func (i *CommentInteractor) UpdateComment(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request nill"))
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}
	reqdata := in.(*UpdateCommentRequest)
	var request *entity.UpdateCommentRequest

	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	resGet, errGer := jwtGen.GetClientMetadata(ctx)
	if errGer != nil {
		tracing.LogError(sp, errGer)
		return nil, errGer
	}

	tracing.LogObject(sp, "GetClientMetadata", resGet)

	request.Username = resGet.Username
	resData, err := i.repo.UpdateComment(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var res *UpdateCommentResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	tracing.LogResponse(sp, res)
	return i.out.UpdateCommentResponse(res)
}
func (i *CommentInteractor) DeleteComment(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request nill"))
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}

	reqdata := in.(*DeleteCommentRequest)
	var request *entity.DeleteCommentRequest

	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	resGet, errGer := jwtGen.GetClientMetadata(ctx)
	if errGer != nil {
		tracing.LogError(sp, errGer)
		return nil, errGer
	}

	tracing.LogObject(sp, "GetClientMetadata", resGet)

	resData, err := i.repo.DeleteComment(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	var res *DeleteCommentResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	tracing.LogResponse(sp, res)
	return i.out.DeleteCommentResponse(res)
}
