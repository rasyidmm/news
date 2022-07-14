package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	service "news/pkg/infrastructure/restful/service/comment"
	"news/pkg/shared/jwtGen"
)

type CommentRouter struct {
	validator *validator.Validate
}

func NewCommentRouter(e *echo.Echo, commentService *service.CommentService) *echo.Echo {
	e.Validator = &UserRouter{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := e.Group("/comment")
	r.POST("/", commentService.CreateComment, jwtGen.IsLoggedIn)
	r.POST("/getall", commentService.GetAllComment, jwtGen.IsLoggedIn)
	r.POST("/getbyall", commentService.GetByAllComment, jwtGen.IsLoggedIn)
	r.POST("/update", commentService.DeleteComment, jwtGen.IsLoggedIn)
	r.POST("/delete", commentService.UpdateComment, jwtGen.IsLoggedIn)
	return e
}
