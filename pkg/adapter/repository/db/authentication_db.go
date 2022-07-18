package db

import (
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"net/http"
	"news/pkg/adapter/db/model"
	"news/pkg/domain/entity"
	"news/pkg/shared/tracing"
	"news/pkg/shared/util"
	"time"
)

type AuthenticationDataHandler struct {
	db *gorm.DB
}

func NewAuthenticationDataHandler(db *gorm.DB) *AuthenticationDataHandler {
	return &AuthenticationDataHandler{db: db}
}

func (d *AuthenticationDataHandler) CheckUsernamePassword(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "CheckUsernamePassword")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.CheckUsernamePasswordRequest)
	var data model.UserModel

	q := d.db.Debug().Where("username = ?", reqData.Username).First(&data)
	if q.Error != nil {
		tracing.LogError(sp, q.Error)
		return nil, status.Error(http.StatusUnauthorized, q.Error.Error())
	}

	tracing.LogObject(sp, "data", data)
	if data.Password != util.HashSha512(reqData.Password) {
		return nil, status.Error(http.StatusUnauthorized, "Authentication Failed")
	}
	res := &entity.CheckUsernamePasswordResponse{
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
		UserId:         data.Id,
	}

	return res, nil
}

func (d *AuthenticationDataHandler) LoginHistorySave(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "LoginHistorySave")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.LoginHistorySaveRequest)

	data := &model.LoginHistoryModel{
		BaseModels: model.BaseModels{CreateBy: reqData.Username},
		BaseCUModels: model.BaseCUModels{
			CreatedAt: time.Now(),
			UpdatedAt: time.Time{},
		},
		Username:    reqData.Username,
		UserId:      reqData.UserId,
		Email:       reqData.Email,
		JenisUser:   reqData.JenisUser,
		ExpiredTime: reqData.ExpiredTime,
		IpAddress:   reqData.IpAddress,
	}

	err := d.db.Debug().Create(&data).Error
	if err != nil {
		tracing.LogError(sp, status.Error(http.StatusBadRequest, err.Error()))
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}
	res := &entity.LoginHistorySaveResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
