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
	CaptionImage string `json:"caption_image"`
	Description  string `json:"description"`
	Editor       string `json:"editor"`
	Tag          string `json:"tag"`
	PublishDate  string `json:"publish_date"`
}

type GetAllNewsRequest struct {
	CurPage int `json:"cur_page"`
	Limit   int `json:"limit"`
}
type GetByAllNewsRequest struct {
	CurPage     int    `json:"cur_page"`
	Limit       int    `json:"limit"`
	Url         string `json:"url"`
	ID          string `json:"id"`
	Title       string `json:"title"`
	Kategori    string `json:"kategori"`
	Description string `json:"description"`
	User        string `json:"user"`
	Editor      string `json:"editor"`
	Tag         string `json:"tag"`
	PublishDate string `json:"publish_date"`
}

type UpdateNewsRequest struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Label        string `json:"label"`
	Url          string `json:"url"`
	Kategori     string `json:"kategori"`
	FeatureImage string `json:"feature_image"`
	CaptionImage string `json:"caption_image"`
	Description  string `json:"description"`
	User         string `json:"user"`
	Editor       string `json:"editor"`
	Tag          string `json:"tag"`
	PublishDate  string `json:"publish_date"`
}

type DeleteNewsRequest struct {
	ID string `json:"id"`
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

func (s *NewsService) GetAllNews(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "GetAllNews")
	defer sp.Finish()

	reqdata := new(GetAllNewsRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.GetAllNewsRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.GetAllNews(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (s *NewsService) GetByAllNew(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "GetByAllNew")
	defer sp.Finish()

	reqdata := new(GetByAllNewsRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.GetByAllNewsRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.GetByAllNews(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (s *NewsService) UpdateNews(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "UpdateNews")
	defer sp.Finish()

	reqdata := new(UpdateNewsRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.UpdateNewsRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.UpdateNews(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (s *NewsService) DeleteNews(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "UpdateNews")
	defer sp.Finish()

	reqdata := new(DeleteNewsRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.DeleteNewsRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.DeleteNews(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)

}
