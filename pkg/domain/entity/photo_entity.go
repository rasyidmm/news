package entity

import "news/pkg/shared/util"

type CreatePhotoRequest struct {
	FileName    string
	TypeFile    string
	PathFile    string
	Description string
}

type CreatePhotoResponse struct {
	StatusCode string
	StatusDesc string
}

type GetPhotoRequest struct {
	CurPage   int
	Limit     int
	NamePhoto string
}

type Photo struct {
	Id          string
	FileName    string
	TypeFile    string
	PathFile    string
	Description string
}

type GetPhotoResponse struct {
	Data             []Photo
	PaginationHelper util.PaginationHelper
}

type GetPhotoByIdRequest struct {
	Id string
}

type GetPhotoByIdResponse struct {
	Id          string
	FileName    string
	TypeFile    string
	PathFile    string
	Description string
}

type UpdatePhotoRequest struct {
	Id          string
	FileName    string
	TypeFile    string
	PathFile    string
	Description string
}
type UpdatePhotoResponse struct {
	StatusCode string
	StatusDesc string
}

type DeletePhotoRequest struct {
	Id string
}
type DeletePhotoResponse struct {
	StatusCode string
	StatusDesc string
}
