package repository

import "github.com/opentracing/opentracing-go"

type KategoriRepository interface {
	KategoriCreate(opentracing.Span, interface{}) (interface{}, error)
	KategoriList(opentracing.Span, interface{}) (interface{}, error)
	KategoriGetById(opentracing.Span, interface{}) (interface{}, error)
	KategoriUpdate(opentracing.Span, interface{}) (interface{}, error)
	KategoriDelete(opentracing.Span, interface{}) (interface{}, error)
}
