package entity

import "news/pkg/shared/util"

type KategoriCreateRequest struct {
	KategoriNama string
	KategoriCode string
}

type KategoriCreateResponse struct {
	StatusCode string
	StatusDesc string
}

type KategoriListRequest struct {
	CurPage    int
	Limit      int
	UserAccess string
}

type Kategory struct {
	Id           string
	KategoriNama string
	KategoriCode string
}

type KategoriListResponse struct {
	Data             []Kategory
	PaginationHelper util.PaginationHelper
}

type KategoriGetByIdRequest struct {
	KategoryId string
}
type KategoriGetByIdResponse struct {
	Id           string
	KategoriNama string
	KategoriCode string
	StatusCode   string
	StatusDesc   string
}

type KategoriUpdateRequest struct {
	Id           string
	KategoriNama string
	KategoriCode string
}
type KategoriUpdateResponse struct {
	StatusCode string
	StatusDesc string
}

type KategoriDeleteRequest struct {
	KategoryId string
}

type KategoriDeleteResponse struct {
	StatusCode string
	StatusDesc string
}
