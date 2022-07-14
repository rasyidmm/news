package model

type NewsModel struct {
	BaseModels
	BaseCUModels
	Title        string `gorm:"column:title;unique"`
	Label        string `gorm:"column:label"`
	Url          string `gorm:"column:url;unique"`
	Kategori     string `gorm:"column:kategori"`
	FeatureImage string `gorm:"column:feature_image"`
	CaptionImage string `gorm:"column:caption_image"`
	Description  string `gorm:"column:description"`
	User         string `gorm:"column:user"`
	Editor       string `gorm:"column:editor"`
	Tag          string `gorm:"column:tag"`
	PublishDate  string `gorm:"column:publish_date"`
}
