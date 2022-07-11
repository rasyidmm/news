package repository

import "github.com/opentracing/opentracing-go"

type AuthenticationRepository interface {
	CheckUsernamePassword(opentracing.Span, interface{}) (interface{}, error)
	LoginHistorySave(opentracing.Span, interface{}) (interface{}, error)
}
