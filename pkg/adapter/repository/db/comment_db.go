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

type CommentDataHandler struct {
	db *gorm.DB
}

func NewCommentDataHandler(db *gorm.DB) *CommentDataHandler {
	return &CommentDataHandler{
		db: db,
	}
}

func (d *CommentDataHandler) CreateComment(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "CreateComment")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	dataReq := in.(*entity.CreateCommentRequest)
	data := &model.CommentModel{
		BaseModels: model.BaseModels{
			CreateBy: dataReq.Username,
		},
		BaseCUModels: model.BaseCUModels{
			CreatedAt: time.Now(),
			UpdatedAt: time.Time{},
		},
		NewsId:      dataReq.NewsId,
		Username:    dataReq.Username,
		CommentText: dataReq.CommentText,
	}

	err := d.db.Debug().Create(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.CreateCommentResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)

	return res, nil
}

func (d *CommentDataHandler) GetAllComment(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetAllNews")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.GetAllCommentRequest)

	var data []model.CommentModel

	q := d.db.Debug()
	q = q.Order("updated_at desc")
	var countAll int64
	q.Model(&model.CommentModel{}).Count(&countAll)
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
	var entityComment []entity.Comment
	for _, item := range data {
		entityComment = append(entityComment, entity.Comment{
			Id:          item.Id,
			NewsId:      item.NewsId,
			Username:    item.Username,
			CommentText: item.CommentText,
		})
	}
	res := &entity.GetAllCommentResponse{
		Data:             entityComment,
		PaginationHelper: *pagination,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
func (d *CommentDataHandler) GetByAllComment(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetAllNews")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.GetByAllCommentRequest)

	var data []model.CommentModel

	q := d.db.Debug()
	q = q.Order("updated_at desc")
	var countAll int64
	q.Model(&model.CommentModel{}).Count(&countAll)
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
	var entityComment []entity.Comment
	for _, item := range data {
		entityComment = append(entityComment, entity.Comment{
			Id:          item.Id,
			NewsId:      item.NewsId,
			Username:    item.Username,
			CommentText: item.CommentText,
		})
	}
	res := &entity.GetByAllCommentResponse{
		Data:             entityComment,
		PaginationHelper: *pagination,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
func (d *CommentDataHandler) UpdateComment(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "UpdateComment")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	dataReq := in.(*entity.UpdateCommentRequest)

	data := model.CommentModel{
		BaseModels:   model.BaseModels{Id: dataReq.Id},
		BaseCUModels: model.BaseCUModels{UpdatedAt: time.Now()},
		Username:     dataReq.Username,
		CommentText:  dataReq.CommentText,
	}
	err := d.db.Debug().Where("id = ?", dataReq.Id).Updates(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.UpdateCommentResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
func (d *CommentDataHandler) DeleteComment(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "DeleteComment")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.DeleteCommentRequest)
	err := d.db.Debug().Where("id = ?", reqData.Id).Delete(&model.NewsModel{}).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.DeleteCommentResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
