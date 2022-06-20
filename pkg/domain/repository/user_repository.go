package repository

import "github.com/opentracing/opentracing-go"

type UserRepository interface {
	UserCreate(opentracing.Span, interface{}) (interface{}, error)
	UserList(opentracing.Span, interface{}) (interface{}, error)
	UserGetById(opentracing.Span, interface{}) (interface{}, error)
	UserUpdate(opentracing.Span, interface{}) (interface{}, error)
	UserDelete(opentracing.Span, interface{}) (interface{}, error)
}
