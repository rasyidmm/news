package entity

type NewsCreateRequest struct {
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
	UserAccess   string
}

type NewsCreateResponse struct {
	StatusCode string
	StatusDesc string
}
