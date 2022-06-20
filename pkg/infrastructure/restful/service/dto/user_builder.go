package dto

import (
	"github.com/mitchellh/mapstructure"
	"news/pkg/shared/util"
)

type UserCreateResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
}
type User struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Twitter        string `json:"twitter"`
	Facebook       string `json:"facebook"`
	Instagram      string `json:"instagram"`
	Biography      string `json:"biography"`
	Email          string `json:"email"`
	NomerHandphone string `json:"nomer_handphone"`
	Password       string `json:"password"`
	JenisUser      string `json:"jenis_user"`
}

type UserListResponse struct {
	Data             []User                `json:"data"`
	PaginationHelper util.PaginationHelper `json:"pagination_helper"`
}

type UserGetByIdResponse struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Twitter        string `json:"twitter"`
	Facebook       string `json:"facebook"`
	Instagram      string `json:"instagram"`
	Biography      string `json:"biography"`
	Email          string `json:"email"`
	NomerHandphone string `json:"nomer_handphone"`
	Password       string `json:"password"`
	JenisUser      string `json:"jenis_user"`
	StatusCode     string `json:"status_code"`
	StatusDesc     string `json:"status_desc"`
}

type UserUpdateResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
}

type UserDeleteResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
}

type UserBuilder struct {
}

func (b *UserBuilder) UserCreateResponse(in interface{}) (interface{}, error) {
	var out *UserCreateResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil

}
func (b *UserBuilder) UserListResponse(in interface{}) (interface{}, error) {
	var out *UserListResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil

}
func (b *UserBuilder) UserGetByIdResponse(in interface{}) (interface{}, error) {
	var out *UserGetByIdResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil

}
func (b *UserBuilder) UserUpdateResponse(in interface{}) (interface{}, error) {
	var out *UserUpdateResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil

}
func (b *UserBuilder) UserDeleteResponse(in interface{}) (interface{}, error) {
	var out *UserDeleteResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil

}
