package model

type NewsModel struct {
	BaseModels
	BaseCUModels
	Title        string `json:"title"`
	Label        string `json:"label"`
	Url          string `json:"url"`
	Kategori     string `json:"kategori"`
	FeatureImage string `json:"feature_image"`
	CationImage  string `json:"cation_image"`
	Description  string `json:"description"`
	User         string `json:"user"`
	Editor       string `json:"editor"`
	Tag          string `json:"tag"`
	PublishDate  string `json:"publish_date"`
}
