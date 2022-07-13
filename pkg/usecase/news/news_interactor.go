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
)

type NewsCreateRequest struct {
	Title        string
	Label        string
	Url          string
	Kategori     string
	FeatureImage string
	CationImage  string
	Description  string
	User         string
	Editor       string
	Tag          string
	PublishDate  string
}

type NewsCreateResponse struct {
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

	resGet := jwtGen.GetClientMetadata(ctx)
	tracing.LogObject(sp, "GetClientMetadata", resGet)
	request.UserAccess = resGet.Username

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
