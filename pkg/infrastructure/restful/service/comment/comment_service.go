package comment

import (
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"news/pkg/shared/tracing"
	usecase "news/pkg/usecase/comment"
)

type CreateCommentRequest struct {
	NewsId      string `json:"news_id"`
	CommentText string `json:"comment_text"`
}

type GetAllCommentRequest struct {
	CurPage int
	Limit   int
}
type GetByAllCommentRequest struct {
	CurPage     int
	Limit       int
	NewsId      string
	Username    string
	CommentText string
}
type UpdateCommentRequest struct {
	Id          string
	Username    string
	CommentText string
}
type DeleteCommentRequest struct {
	Id string
}

type CommentService struct {
	uc usecase.CommentInputPort
}

func NewCommentService(u usecase.CommentInputPort) *CommentService {
	return &CommentService{
		uc: u,
	}
}

func (s *CommentService) CreateComment(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "CreateComment")
	defer sp.Finish()

	reqdata := new(CreateCommentRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.CreateCommentRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.CreateComment(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (s *CommentService) GetAllComment(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "GetAllComment")
	defer sp.Finish()

	reqdata := new(GetAllCommentRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.GetAllCommentRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.GetAllComment(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}
func (s *CommentService) GetByAllComment(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "GetByAllComment")
	defer sp.Finish()

	reqdata := new(GetByAllCommentRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.GetByAllCommentRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.GetByAllComment(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}
func (s *CommentService) UpdateComment(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "UpdateComment")
	defer sp.Finish()

	reqdata := new(UpdateCommentRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.UpdateCommentRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.UpdateComment(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}
func (s *CommentService) DeleteComment(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "DeleteComment")
	defer sp.Finish()

	reqdata := new(DeleteCommentRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.DeleteCommentRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.DeleteComment(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}
