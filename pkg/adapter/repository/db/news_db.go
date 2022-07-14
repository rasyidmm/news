package db

import (
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"news/pkg/adapter/db/model"
	"news/pkg/domain/entity"
	"news/pkg/shared/tracing"
	"news/pkg/shared/util"
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
		CaptionImage: dataReq.CaptionImage,
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

func (d *NewsDataHandle) GetAllNews(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetAllNews")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.GetAllNewsRequest)

	var data []model.NewsModel

	q := d.db.Debug()
	q = q.Order("updated_at desc")
	var countAll int64
	q.Model(&model.NewsModel{}).Count(&countAll)
	pagination := util.MakePagination(reqData.CurPage, reqData.Limit, int(countAll))

	q = q.Offset(pagination.FormIndex - 1)
	if reqData.Limit != 0 {
		q = q.Limit(pagination.PerPage)
	}
	err := q.Find(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	var entityNews []entity.News
	for _, item := range data {
		entityNews = append(entityNews, entity.News{
			ID:           item.Id,
			Title:        item.Title,
			Label:        item.Label,
			Url:          item.Url,
			Kategori:     item.Kategori,
			FeatureImage: item.FeatureImage,
			CaptionImage: item.CaptionImage,
			Description:  item.Description,
			User:         item.User,
			Editor:       item.Editor,
			Tag:          item.Tag,
			PublishDate:  item.PublishDate,
		})
	}
	res := &entity.GetAllNewsResponse{
		Data:             entityNews,
		PaginationHelper: *pagination,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
func (d *NewsDataHandle) GetByAllNews(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetByAllNews")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.GetByAllNewsRequest)

	var data []model.NewsModel

	q := d.db.Debug()

	if reqData.Url != "" {
		q = q.Where("url = ?", reqData.Url)
	}
	if reqData.Title != "" {
		q = q.Where("title = ?", reqData.Title)
	}
	if reqData.Kategori != "" {
		q = q.Where("kategori = ?", reqData.Kategori)
	}
	if reqData.Description != "" {
		descrip := "%" + reqData.Description + "%"
		q = q.Where("description like ?", descrip)
	}
	if reqData.User != "" {
		q = q.Where("username = ?", reqData.User)
	}
	if reqData.Editor != "" {
		q = q.Where("editor = ?", reqData.Editor)
	}
	if reqData.Tag != "" {
		q = q.Where("tag = ?", reqData.Tag)
	}
	if reqData.PublishDate != "" {
		datenew := reqData.PublishDate + "%"
		q = q.Where("publish_date like ?", datenew)
	}

	q = q.Order("updated_at desc")

	var countAll int64
	q.Model(&model.NewsModel{}).Count(&countAll)
	pagination := util.MakePagination(reqData.CurPage, reqData.Limit, int(countAll))

	q = q.Offset(pagination.FormIndex - 1)
	if reqData.Limit != 0 {
		q = q.Limit(pagination.PerPage)
	}
	err := q.Find(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	var entityNews []entity.News
	for _, item := range data {
		entityNews = append(entityNews, entity.News{
			ID:           item.Id,
			Title:        item.Title,
			Label:        item.Label,
			Url:          item.Url,
			Kategori:     item.Kategori,
			FeatureImage: item.FeatureImage,
			CaptionImage: item.CaptionImage,
			Description:  item.Description,
			User:         item.User,
			Editor:       item.Editor,
			Tag:          item.Tag,
			PublishDate:  item.PublishDate,
		})
	}
	res := &entity.GetByAllNewsResponse{
		Data:             entityNews,
		PaginationHelper: *pagination,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *NewsDataHandle) UpdateNews(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "UpdateNews")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	dataReq := in.(*entity.UpdateNewsRequest)

	data := model.NewsModel{
		BaseModels:   model.BaseModels{Id: dataReq.ID},
		BaseCUModels: model.BaseCUModels{UpdatedAt: time.Now()},
		Title:        dataReq.Title,
		Label:        dataReq.Label,
		Url:          dataReq.Url,
		Kategori:     dataReq.Kategori,
		FeatureImage: dataReq.FeatureImage,
		CaptionImage: dataReq.CaptionImage,
		Description:  dataReq.Description,
		User:         dataReq.User,
		Editor:       dataReq.Editor,
		Tag:          dataReq.Tag,
		PublishDate:  dataReq.PublishDate,
	}
	err := d.db.Debug().Where("id = ?", dataReq.ID).Updates(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.KategoriUpdateResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *NewsDataHandle) DeleteNews(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "DeleteNews")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.DeleteNewsRequest)
	err := d.db.Debug().Where("id = ?", reqData.ID).Delete(&model.NewsModel{}).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.DeleteNewsResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
