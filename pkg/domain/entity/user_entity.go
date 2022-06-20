package entity

import "news/pkg/shared/util"

type UserCreateRequest struct {
	Username       string
	FirstName      string
	LastName       string
	Twitter        string
	Facebook       string
	Instagram      string
	Biography      string
	Email          string
	NomerHandphone string
	Password       string
	JenisUser      string
}
type UserCreateResponse struct {
	StatusCode string
	StatusDesc string
}
type UserListRequest struct {
	CurPage int
	Limit   int
}
type User struct {
	Id             string
	Username       string
	FirstName      string
	LastName       string
	Twitter        string
	Facebook       string
	Instagram      string
	Biography      string
	Email          string
	NomerHandphone string
	Password       string
	JenisUser      string
}

type UserListResponse struct {
	Data             []User
	PaginationHelper util.PaginationHelper
}
type UserGetByIdRequest struct {
	UserId string
}
type UserGetByIdResponse struct {
	Id             string
	Username       string
	FirstName      string
	LastName       string
	Twitter        string
	Facebook       string
	Instagram      string
	Biography      string
	Email          string
	NomerHandphone string
	Password       string
	JenisUser      string
	StatusCode     string
	StatusDesc     string
}
type UserUpdateRequest struct {
	Id             string
	Username       string
	FirstName      string
	LastName       string
	Twitter        string
	Facebook       string
	Instagram      string
	Biography      string
	Email          string
	NomerHandphone string
	Password       string
	JenisUser      string
}
type UserUpdateResponse struct {
	StatusCode string
	StatusDesc string
}

type UserDeleteRequest struct {
	UserId string
}

type UserDeleteResponse struct {
	StatusCode string
	StatusDesc string
}
