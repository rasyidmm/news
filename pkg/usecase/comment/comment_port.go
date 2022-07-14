package comment

import "github.com/labstack/echo/v4"

type CommentInputPort interface {
	CreateComment(ctx echo.Context, in interface{}) (interface{}, error)
	GetAllComment(ctx echo.Context, in interface{}) (interface{}, error)
	GetByAllComment(ctx echo.Context, in interface{}) (interface{}, error)
	UpdateComment(ctx echo.Context, in interface{}) (interface{}, error)
	DeleteComment(ctx echo.Context, in interface{}) (interface{}, error)
}

type CommentOutputPort interface {
	CreateCommentResponse(interface{}) (interface{}, error)
	GetAllCommentResponse(interface{}) (interface{}, error)
	GetByAllCommentResponse(interface{}) (interface{}, error)
	UpdateCommentResponse(interface{}) (interface{}, error)
	DeleteCommentResponse(interface{}) (interface{}, error)
}
