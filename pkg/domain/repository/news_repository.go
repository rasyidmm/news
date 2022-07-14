package repository

import (
	"github.com/opentracing/opentracing-go"
)

type NewsRepository interface {
	CreateNews(opentracing.Span, interface{}) (interface{}, error)
	GetAllNews(opentracing.Span, interface{}) (interface{}, error)
	GetByAllNews(opentracing.Span, interface{}) (interface{}, error)
	UpdateNews(opentracing.Span, interface{}) (interface{}, error)
	DeleteNews(opentracing.Span, interface{}) (interface{}, error)
}
