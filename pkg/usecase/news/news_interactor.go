package news

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

type NewsCreateRequest struct {
	Title        string
	Label        string
	Url          string
	Kategori     string
	FeatureImage string
	CaptionImage string
	Description  string
	Editor       string
	Tag          string
	PublishDate  string
}

type NewsCreateResponse struct {
	StatusCode string
	StatusDesc string
}

type News struct {
	ID           string
	Title        string
	Label        string
	Url          string
	Kategori     string
	FeatureImage string
	CaptionImage string
	Description  string
	User         string
	Editor       string
	Tag          string
	PublishDate  string
}

type GetAllNewsRequest struct {
	CurPage int
	Limit   int
}

type GetAllNewsResponse struct {
	Data             []News
	PaginationHelper util.PaginationHelper
}

type GetByAllNewsRequest struct {
	CurPage     int
	Limit       int
	Url         string
	ID          string
	Title       string
	Kategori    string
	Description string
	User        string
	Editor      string
	Tag         string
	PublishDate string
}

type GetByAllNewsResponse struct {
	Data             []News
	PaginationHelper util.PaginationHelper
}

type UpdateNewsRequest struct {
	ID           string
	Title        string
	Label        string
	Url          string
	Kategori     string
	FeatureImage string
	CaptionImage string
	Description  string
	User         string
	Editor       string
	Tag          string
	PublishDate  string
}

type UpdateNewsResponse struct {
	StatusCode string
	StatusDesc string
}

type DeleteNewsRequest struct {
	ID string
}

type DeleteNewsResponse struct {
	StatusCode string
	StatusDesc string
}

type NewsInteractor struct {
	repo repository.NewsRepository
	out  NewsOutputPort
}

func NewNewsInteractor(r repository.NewsRepository, o NewsOutputPort) *NewsInteractor {
	return &NewsInteractor{
		repo: r,
		out:  o,
	}
}

func (i *NewsInteractor) CreateNews(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request nill"))
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}

	reqdata := in.(*NewsCreateRequest)
	var request *entity.NewsCreateRequest

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
	request.UserAccess = resGet.Username
	request.User = resGet.Username

	resData, err := i.repo.CreateNews(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var res *NewsCreateResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	tracing.LogResponse(sp, res)
	return i.out.CreateNewsResponse(res)

}

func (i *NewsInteractor) GetAllNews(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request nill"))
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}

	reqdata := in.(*GetAllNewsRequest)
	var request *entity.GetAllNewsRequest

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

	resData, err := i.repo.GetAllNews(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var res *GetAllNewsResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	tracing.LogResponse(sp, res)
	return i.out.GetAllNewsResponse(res)

}
func (i *NewsInteractor) GetByAllNews(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request nill"))
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}

	reqdata := in.(*GetByAllNewsRequest)
	var request *entity.GetByAllNewsRequest

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

	resData, err := i.repo.GetByAllNews(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var res *GetByAllNewsResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	tracing.LogResponse(sp, res)
	return i.out.GetByAllNewResponse(res)

}
func (i *NewsInteractor) UpdateNews(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request nill"))
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}

	reqdata := in.(*UpdateNewsRequest)
	var request *entity.UpdateNewsRequest

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

	resData, err := i.repo.UpdateNews(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var res *UpdateNewsResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	tracing.LogResponse(sp, res)
	return i.out.UpdateNewsResponse(res)
}

func (i *NewsInteractor) DeleteNews(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request nill"))
		return nil, status.Error(codes.InvalidArgument, "request nill")
	}

	reqdata := in.(*DeleteNewsRequest)
	var request *entity.DeleteNewsRequest

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

	resData, err := i.repo.DeleteNews(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var res *DeleteNewsResponse

	err = mapstructure.Decode(resData, &res)
	if err != nil {
		tracing.LogError(sp, status.Error(codes.InvalidArgument, "request parsing err"))
		return nil, status.Error(codes.InvalidArgument, "request parsing err")
	}

	tracing.LogResponse(sp, res)
	return i.out.DeleteNewsResponse(res)
}
