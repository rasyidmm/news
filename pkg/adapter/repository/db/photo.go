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

type PhotoDataHandler struct {
	db *gorm.DB
}

func NewPhotoDataHanlder(db *gorm.DB) *PhotoDataHandler {
	return &PhotoDataHandler{
		db: db,
	}
}

func (d *PhotoDataHandler) PhotoCreate(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "PhotoCreate")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	dataReq := in.(*entity.CreatePhotoRequest)
	data := &model.PhotoModel{
		BaseModels: model.BaseModels{},
		BaseCUModels: model.BaseCUModels{
			CreatedAt: time.Now(),
			UpdatedAt: time.Time{},
		},
		TypeFile:    dataReq.TypeFile,
		PathFile:    dataReq.PathFile,
		FileName:    dataReq.FileName,
		Description: dataReq.Description,
	}
	err := d.db.Debug().Create(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	res := &entity.CreatePhotoResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
func (d *PhotoDataHandler) PhotoGet(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "PhotoGet")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.GetPhotoRequest)

	var data []model.PhotoModel

	q := d.db.Debug()
	if reqData.NamePhoto != "" {
		name := "%" + reqData.NamePhoto + "%"
		q = q.Where("file_name like ? ", name)
	}

	q = q.Order("id desc")
	var countAll int64
	q.Model(&model.PhotoModel{}).Count(&countAll)
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
	var entityPhoto []entity.Photo
	for _, item := range data {
		entityPhoto = append(entityPhoto, entity.Photo{
			Id:          item.Id,
			FileName:    item.FileName,
			TypeFile:    item.TypeFile,
			PathFile:    item.PathFile,
			Description: item.Description,
		})
	}
	res := &entity.GetPhotoResponse{
		Data:             entityPhoto,
		PaginationHelper: *pagination,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *PhotoDataHandler) PhotoGetById(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "PhotoGetById")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.GetPhotoByIdRequest)

	var data model.PhotoModel
	err := d.db.Debug().Where("id = ?", reqData.Id).First(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.GetPhotoByIdResponse{
		Id:          data.Id,
		FileName:    data.FileName,
		TypeFile:    data.TypeFile,
		PathFile:    data.PathFile,
		Description: data.Description,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *PhotoDataHandler) PhotoUpdate(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "PhotoGetById")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.UpdatePhotoRequest)

	data := &model.PhotoModel{
		BaseModels:   model.BaseModels{Id: reqData.Id},
		BaseCUModels: model.BaseCUModels{UpdatedAt: time.Now()},
		FileName:     reqData.FileName,
		TypeFile:     reqData.TypeFile,
		PathFile:     reqData.PathFile,
		Description:  reqData.Description,
	}
	err := d.db.Debug().Where("id = ?", reqData.Id).Updates(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.UpdatePhotoResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *PhotoDataHandler) PhotoDelete(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "PhotoDelete")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.DeletePhotoRequest)
	err := d.db.Debug().Where("id = ?", reqData.Id).Delete(&model.PhotoModel{}).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.DeletePhotoResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
