package dto

import (
	"github.com/mitchellh/mapstructure"
	"news/pkg/shared/util"
)

type CreatePhotoResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
}

type Photo struct {
	Id          string `json:"id"`
	FileBase64  string `json:"file_base_64"`
	FileName    string `json:"file_name"`
	TypeFile    string `json:"type_file"`
	Description string `json:"description"`
}

type PhotoListResponse struct {
	Data             []Photo               `json:"data"`
	PaginationHelper util.PaginationHelper `json:"pagination_helper"`
}

type UpdatePhotoResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
}

type PhotoGetByIdResponse struct {
	Id          string `json:"id"`
	FileBase64  string `json:"file_base_64"`
	FileName    string `json:"file_name"`
	TypeFile    string `json:"type_file"`
	Description string `json:"description"`
}

type DeletePhotoResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
}

type PhotoBuilder struct{}

func (b *PhotoBuilder) PhotoCreateResponse(in interface{}) (interface{}, error) {
	var out *CreatePhotoResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *PhotoBuilder) PhotoGetResponse(in interface{}) (interface{}, error) {
	var out *PhotoListResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *PhotoBuilder) PhotoUpdateResponse(in interface{}) (interface{}, error) {
	var out *UpdatePhotoResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *PhotoBuilder) PhotoGetByIdResponse(in interface{}) (interface{}, error) {
	var out *PhotoGetByIdResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (b *PhotoBuilder) PhotoDeleteByIdResponse(in interface{}) (interface{}, error) {
	var out *DeletePhotoResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
