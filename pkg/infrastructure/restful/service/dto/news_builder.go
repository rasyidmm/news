package dto

import (
	"github.com/mitchellh/mapstructure"
	"news/pkg/shared/util"
)

type News struct {
	ID           string
	Title        string
	Label        string
	Url          string
	Kategori     string
	FeatureImage string
	CationImage  string
	Description  string
	User         string
	Editor       string
	Tag          string
	PublishDate  string
}

type NewsCreateResponse struct {
	StatusCode string
	StatusDesc string
}

type GetAllNewResponse struct {
	Data             []News
	PaginationHelper util.PaginationHelper
}

type GetByAllNewResponse struct {
	Data             []News
	PaginationHelper util.PaginationHelper
}

type UpdateNewsResponse struct {
	StatusCode string
	StatusDesc string
}

type DeleteNewsResponse struct {
	StatusCode string
	StatusDesc string
}
type NewsBuilder struct {
}

func (b *NewsBuilder) CreateNewsResponse(in interface{}) (interface{}, error) {
	var out *NewsCreateResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
	return nil, nil
}

func (b *NewsBuilder) GetAllNewsResponse(in interface{}) (interface{}, error) {
	var out *GetAllNewResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (b *NewsBuilder) GetByAllNewResponse(in interface{}) (interface{}, error) {
	var out *GetByAllNewResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *NewsBuilder) UpdateNewsResponse(in interface{}) (interface{}, error) {
	var out *UpdateNewsResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *NewsBuilder) DeleteNewsResponse(in interface{}) (interface{}, error) {
	var out *DeleteNewsResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
