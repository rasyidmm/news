package dto

import (
	"github.com/mitchellh/mapstructure"
	"news/pkg/shared/util"
)

type KategoriCreateResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
}
type Kategory struct {
	Id           string `json:"id"`
	KategoriNama string `json:"kategorinama"`
	KategoriCode string `json:"kategoricode"`
}
type KategoriListResponse struct {
	Data             []Kategory            `json:"data"`
	PaginationHelper util.PaginationHelper `json:"paginationHelper"`
}
type KategoriGetByIdResponse struct {
	Id           string `json:"id"`
	KategoriNama string `json:"kategorinama"`
	KategoriCode string `json:"kategoricode"`
	StatusCode   string `json:"status_code"`
	StatusDesc   string `json:"status_desc"`
}
type KategoriUpdateResponse struct {
	StatusCode string `json:"statusCode"`
	StatusDesc string `json:"statusDesc"`
}
type KategoriDeleteResponse struct {
	StatusCode string `json:"statusCode"`
	StatusDesc string `json:"statusDesc"`
}

type KategoriBuilder struct {
}

func (b *KategoriBuilder) KategoriCreateResponse(in interface{}) (interface{}, error) {
	var out *KategoriCreateResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (b *KategoriBuilder) KategoriListResponse(in interface{}) (interface{}, error) {
	var out *KategoriListResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (b *KategoriBuilder) KategoriGetByIdResponse(in interface{}) (interface{}, error) {
	var out *KategoriGetByIdResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (b *KategoriBuilder) KategoriUpdateResponse(in interface{}) (interface{}, error) {
	var out *KategoriUpdateResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (b *KategoriBuilder) KategoriDeleteResponse(in interface{}) (interface{}, error) {
	var out *KategoriDeleteResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
