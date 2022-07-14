package entity

import "news/pkg/shared/util"

type NewsCreateRequest struct {
	Title        string
	Label        string
	Url          string
	Kategori     string
	FeatureImage string
	CaptionImage string
	Description  string
	User         string
	Editor       string
	Tag          string
	PublishDate  string
	UserAccess   string
}

type NewsCreateResponse struct {
	StatusCode string
	StatusDesc string
}

type News struct {
	ID           string
	Title        string
	Label        string
	Url          string
	Kategori     string
	FeatureImage string
	CaptionImage string
	Description  string
	User         string
	Editor       string
	Tag          string
	PublishDate  string
}

type GetAllNewsRequest struct {
	CurPage int
	Limit   int
}

type GetAllNewsResponse struct {
	Data             []News
	PaginationHelper util.PaginationHelper
}

type GetByAllNewsRequest struct {
	CurPage     int
	Limit       int
	Url         string
	Title       string
	Kategori    string
	Description string
	User        string
	Editor      string
	Tag         string
	PublishDate string
}

type GetByAllNewsResponse struct {
	Data             []News
	PaginationHelper util.PaginationHelper
}

type UpdateNewsRequest struct {
	ID           string
	Title        string
	Label        string
	Url          string
	Kategori     string
	FeatureImage string
	CaptionImage string
	Description  string
	User         string
	Editor       string
	Tag          string
	PublishDate  string
}

type UpdateNewsResponse struct {
	StatusCode string
	StatusDesc string
}

type DeleteNewsRequest struct {
	ID string
}

type DeleteNewsResponse struct {
	StatusCode string
	StatusDesc string
}
