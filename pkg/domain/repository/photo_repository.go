package repository

import "github.com/opentracing/opentracing-go"

type PhotoRepository interface {
	PhotoCreate(opentracing.Span, interface{}) (interface{}, error)
	PhotoGet(opentracing.Span, interface{}) (interface{}, error)
	PhotoGetById(opentracing.Span, interface{}) (interface{}, error)
	PhotoUpdate(opentracing.Span, interface{}) (interface{}, error)
	PhotoDelete(opentracing.Span, interface{}) (interface{}, error)
}
