package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModels struct to generate model id
type BaseModels struct {
	Id       string `json:"id" gorm:"primaryKey"`
	CreateBy string `gorm:"create_by"`
	UpdateBy string `gorm:"update_by"`
}

// BaseCUModels struct to generate CreatedAt, UpdatedAt
type BaseCUModels struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Tabler ...
type Tabler interface {
	TableName() string
}

// BeforeCreate create uuid before model create
func (base *BaseModels) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	base.Id = uuid.String()
	return nil
}
