package dto

import "github.com/mitchellh/mapstructure"

type NewsCreateResponse struct {
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
