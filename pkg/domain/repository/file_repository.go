package repository

import "github.com/opentracing/opentracing-go"

type FileLocalRepository interface {
	SaveFile(opentracing.Span, interface{}) (interface{}, error)
	GetFile(opentracing.Span, interface{}) (interface{}, error)
	UpdateFile(opentracing.Span, interface{}) (interface{}, error)
	RemoveFile(opentracing.Span, interface{}) (interface{}, error)
}
