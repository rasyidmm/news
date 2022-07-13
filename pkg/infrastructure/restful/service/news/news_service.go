package news

import (
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"news/pkg/shared/tracing"
	usecase "news/pkg/usecase/news"
)

type NewsCreateRequest struct {
	Title        string `json:"title"`
	Label        string `json:"label"`
	Url          string `json:"url"`
	Kategori     string `json:"kategori"`
	FeatureImage string `json:"feature_image"`
	CationImage  string `json:"cation_image"`
	Description  string `json:"description"`
	User         string `json:"user"`
	Editor       string `json:"editor"`
	Tag          string `json:"tag"`
	PublishDate  string `json:"publish_date"`
}

type NewsService struct {
	uc usecase.NewsInputPort
}

func NewNewsService(u usecase.NewsInputPort) *NewsService {
	return &NewsService{
		uc: u,
	}
}

func (s *NewsService) CreateNews(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "CreateNews")
	defer sp.Finish()

	reqdata := new(NewsCreateRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.NewsCreateRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.CreateNews(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}
