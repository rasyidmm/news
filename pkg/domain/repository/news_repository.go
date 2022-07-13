package repository

import (
	"github.com/opentracing/opentracing-go"
)

type NewsRepository interface {
	CreateNews(opentracing.Span, interface{}) (interface{}, error)
}
