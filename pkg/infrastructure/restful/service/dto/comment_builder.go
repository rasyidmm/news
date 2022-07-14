package dto

import (
	"github.com/mitchellh/mapstructure"
	"news/pkg/shared/util"
)

type CreateCommentResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
}

type Comment struct {
	NewsId      string
	Username    string
	CommentText string
}
type GetAllCommentResponse struct {
	Data             []Comment
	PaginationHelper util.PaginationHelper
}

type GetByAllCommentResponse struct {
	Data             []Comment
	PaginationHelper util.PaginationHelper
}

type UpdateCommentResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
}

type DeleteCommentResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
}

type CommentBuilder struct {
}

func (b *CommentBuilder) CreateCommentResponse(in interface{}) (interface{}, error) {
	var out *CreateCommentResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *CommentBuilder) GetAllCommentResponse(in interface{}) (interface{}, error) {
	var out *GetAllCommentResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *CommentBuilder) GetByAllCommentResponse(in interface{}) (interface{}, error) {
	var out *GetByAllCommentResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *CommentBuilder) UpdateCommentResponse(in interface{}) (interface{}, error) {
	var out *UpdateCommentResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *CommentBuilder) DeleteCommentResponse(in interface{}) (interface{}, error) {
	var out *DeleteCommentResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
