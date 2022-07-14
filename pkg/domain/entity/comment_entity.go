package entity

import "news/pkg/shared/util"

type CreateCommentRequest struct {
	NewsId      string
	Username    string
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
