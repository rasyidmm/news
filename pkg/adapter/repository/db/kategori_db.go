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

type KategoryDatahandler struct {
	db *gorm.DB
}

func NewKategoryDatahandler(db *gorm.DB) *KategoryDatahandler {
	return &KategoryDatahandler{db: db}
}

func (d *KategoryDatahandler) KategoriCreate(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "KategoriCreate")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	dataReq := in.(*entity.KategoriCreateRequest)
	data := &model.KategoriModel{
		BaseModels: model.BaseModels{},
		BaseCUModels: model.BaseCUModels{
			CreatedAt: time.Now(),
			UpdatedAt: time.Time{},
		},
		KategoriName: dataReq.KategoriNama,
		KategoriCode: dataReq.KategoriCode,
	}
	err := d.db.Debug().Create(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	res := &entity.KategoriCreateResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *KategoryDatahandler) KategoriList(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "KategoriList")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.KategoriListRequest)

	var data []model.KategoriModel

	q := d.db.Debug()
	q = q.Order("id desc")
	var countAll int64
	q.Model(&model.KategoriModel{}).Count(&countAll)
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
	var entityKategory []entity.Kategory
	for _, item := range data {
		entityKategory = append(entityKategory, entity.Kategory{
			Id:           item.Id,
			KategoriNama: item.KategoriName,
			KategoriCode: item.KategoriCode,
		})
	}
	res := &entity.KategoriListResponse{
		Data:             entityKategory,
		PaginationHelper: *pagination,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *KategoryDatahandler) KategoriGetById(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "KategoriGetById")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.KategoriGetByIdRequest)

	var data model.KategoriModel
	err := d.db.Debug().Where("id = ?", reqData.KategoryId).First(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	res := &entity.KategoriGetByIdResponse{
		Id:           data.Id,
		KategoriNama: data.KategoriName,
		KategoriCode: data.KategoriCode,
		StatusCode:   "00",
		StatusDesc:   "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *KategoryDatahandler) KategoriUpdate(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "KategoriUpdate")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.KategoriUpdateRequest)

	data := model.KategoriModel{
		BaseModels:   model.BaseModels{Id: reqData.Id},
		BaseCUModels: model.BaseCUModels{UpdatedAt: time.Now()},
		KategoriName: reqData.KategoriNama,
		KategoriCode: reqData.KategoriCode,
	}
	err := d.db.Debug().Where("id = ?", reqData.Id).Updates(&data).Error
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

func (d *KategoryDatahandler) KategoriDelete(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "KategoriDelete")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.KategoriDeleteRequest)
	err := d.db.Debug().Where("id = ?", reqData.KategoryId).Delete(&model.KategoriModel{}).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.KategoriDeleteResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
