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

type UserDataHandler struct {
	db *gorm.DB
}

func NewUserDataHandler(db *gorm.DB) *UserDataHandler {
	return &UserDataHandler{
		db: db,
	}

}

func (d *UserDataHandler) UserCreate(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "UserCreate")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.UserCreateRequest)
	data := &model.UserModel{
		BaseModels: model.BaseModels{},
		BaseCUModels: model.BaseCUModels{
			CreatedAt: time.Now(),
			UpdatedAt: time.Time{},
		},
		Username:       reqdata.Username,
		FirstName:      reqdata.FirstName,
		LastName:       reqdata.LastName,
		Twitter:        reqdata.Twitter,
		Facebook:       reqdata.Facebook,
		Instagram:      reqdata.Instagram,
		Biography:      reqdata.Biography,
		Email:          reqdata.Email,
		NomerHandphone: reqdata.NomerHandphone,
		Password:       reqdata.Password,
		JenisUser:      reqdata.JenisUser,
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

func (d *UserDataHandler) UserList(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "UserList")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.UserListRequest)

	var data []model.UserModel

	q := d.db.Debug()
	q = q.Order("id desc")
	var countAll int64
	q.Model(&model.UserModel{}).Count(&countAll)
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
	var entityUser []entity.User
	for _, reqdata := range data {
		entityUser = append(entityUser, entity.User{
			Id:             reqdata.Id,
			Username:       reqdata.Username,
			FirstName:      reqdata.FirstName,
			LastName:       reqdata.LastName,
			Twitter:        reqdata.Twitter,
			Facebook:       reqdata.Facebook,
			Instagram:      reqdata.Instagram,
			Biography:      reqdata.Biography,
			Email:          reqdata.Email,
			NomerHandphone: reqdata.NomerHandphone,
			Password:       reqdata.Password,
			JenisUser:      reqdata.JenisUser,
		})
	}
	res := &entity.UserListResponse{
		Data:             entityUser,
		PaginationHelper: *pagination,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *UserDataHandler) UserGetById(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "UserGetById")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.UserGetByIdRequest)

	var data model.UserModel
	err := d.db.Debug().Where("id = ?", reqData.UserId).First(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	res := &entity.UserGetByIdResponse{
		Id:             data.Id,
		Username:       data.Username,
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		Twitter:        data.Twitter,
		Facebook:       data.Facebook,
		Instagram:      data.Instagram,
		Biography:      data.Biography,
		Email:          data.Email,
		NomerHandphone: data.NomerHandphone,
		Password:       data.Password,
		JenisUser:      data.JenisUser,
		StatusCode:     "00",
		StatusDesc:     "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *UserDataHandler) UserUpdate(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "UserUpdate")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.UserUpdateRequest)

	data := model.UserModel{
		BaseModels:     model.BaseModels{Id: reqData.Id},
		BaseCUModels:   model.BaseCUModels{UpdatedAt: time.Now()},
		Username:       reqData.Username,
		FirstName:      reqData.FirstName,
		LastName:       reqData.LastName,
		Twitter:        reqData.Twitter,
		Facebook:       reqData.Facebook,
		Instagram:      reqData.Instagram,
		Biography:      reqData.Biography,
		Email:          reqData.Email,
		NomerHandphone: reqData.NomerHandphone,
		Password:       reqData.Password,
		JenisUser:      reqData.JenisUser,
	}
	err := d.db.Debug().Where("id = ?", reqData.Id).Updates(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.UserUpdateResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *UserDataHandler) UserDelete(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "UserDelete")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.UserDeleteRequest)
	err := d.db.Debug().Where("id = ?", reqData.UserId).Delete(&model.UserModel{}).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.UserDeleteResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
