package comment

import (
	"github.com/labstack/echo/v4"
	usecase "news/pkg/usecase/comment"
)

type CommentService struct {
	uc usecase.CommentInputPort
}

func NewCommentService(u usecase.CommentInputPort) *CommentService {
	return &CommentService{
		uc: u,
	}
}

func (s *CommentService) CreateComment(c echo.Context) error {
	return nil
}
