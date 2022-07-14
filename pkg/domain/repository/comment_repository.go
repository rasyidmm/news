package repository

import "github.com/opentracing/opentracing-go"

type CommentRepository interface {
	CreateComment(span opentracing.Span, in interface{}) (interface{}, error)
	GetAllComment(span opentracing.Span, in interface{}) (interface{}, error)
	GetByAllComment(span opentracing.Span, in interface{}) (interface{}, error)
	UpdateComment(span opentracing.Span, in interface{}) (interface{}, error)
	DeleteComment(span opentracing.Span, in interface{}) (interface{}, error)
}
