package db

import (
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"news/pkg/adapter/db/model"
	"news/pkg/domain/entity"
	"news/pkg/shared/tracing"
	"time"
)

type NewsDataHandle struct {
	db *gorm.DB
}

func NewNewsDataHandle(db *gorm.DB) *NewsDataHandle {
	return &NewsDataHandle{
		db: db,
	}
}

func (d *NewsDataHandle) CreateNews(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "CreateNews")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	dataReq := in.(*entity.NewsCreateRequest)
	data := &model.NewsModel{
		BaseModels: model.BaseModels{
			CreateBy: dataReq.UserAccess,
		},
		BaseCUModels: model.BaseCUModels{
			CreatedAt: time.Now(),
			UpdatedAt: time.Time{},
		},
		Title:        dataReq.Title,
		Label:        dataReq.Label,
		Url:          dataReq.Url,
		Kategori:     dataReq.Kategori,
		FeatureImage: dataReq.FeatureImage,
		CationImage:  dataReq.CationImage,
		Description:  dataReq.Description,
		User:         dataReq.User,
		Editor:       dataReq.Editor,
		Tag:          dataReq.Tag,
		PublishDate:  dataReq.PublishDate,
	}
	err := d.db.Debug().Create(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	res := &entity.NewsCreateResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
