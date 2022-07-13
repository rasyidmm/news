package model

type CommentModel struct {
	BaseModels
	BaseCUModels
	NewsId      string `gorm:"column:news_id"`
	Username    string `gorm:"column:username"`
	CommentText string `gorm:"column:comment_text"`
}
