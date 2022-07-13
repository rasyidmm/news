package router

import (
	"github.com/labstack/echo/v4"
	service "news/pkg/infrastructure/restful/service/comment"
)

type CommentRouter struct {
}

func NewCommentRouter(e *echo.Echo, commentService *service.CommentService) *echo.Echo {
	return e
}
