package db

import "gorm.io/gorm"

type CommentDataHandler struct {
	db *gorm.DB
}

func NewCommentDataHandler(db *gorm.DB) *CommentDataHandler {
	return &CommentDataHandler{
		db: db,
	}
}
