package photo

import (
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"
	"net/http"
	"news/pkg/domain/entity"
	"news/pkg/domain/repository"
	"news/pkg/shared/enum"
	"news/pkg/shared/tracing"
	"news/pkg/shared/util"
)

type CreatePhotoRequest struct {
	FileBase64  string
	FileName    string
	TypeFile    string
	Description string
}

type CreatePhotoResponse struct {
	StatusCode string
	StatusDesc string
}

type PhotoListRequest struct {
	CurPage   int
	Limit     int
	NamePhoto string
}

type Photo struct {
	Id          string
	FileBase64  string
	FileName    string
	TypeFile    string
	Description string
}

type PhotoListResponse struct {
	Data             []Photo
	PaginationHelper util.PaginationHelper
}

type UpdatePhotoRequest struct {
	Id          string
	FileBase64  string
	FileName    string
	TypeFile    string
	Description string
}

type UpdatePhotoResponse struct {
	StatusCode string
	StatusDesc string
}

type PhotoGetByIdRequest struct {
	Id string
}
type PhotoGetByIdResponse struct {
	Id          string
	FileBase64  string
	FileName    string
	TypeFile    string
	Description string
}

type DeletePhotoRequest struct {
	Id string
}
type DeletePhotoResponse struct {
	StatusCode string
	StatusDesc string
}

type PhotoInteractor struct {
	repo     repository.PhotoRepository
	repoFile repository.FileLocalRepository
	out      PhotoOutputPort
}

func NewPhotoInteractor(r repository.PhotoRepository, rf repository.FileLocalRepository, o PhotoOutputPort) *PhotoInteractor {
	return &PhotoInteractor{
		repo:     r,
		repoFile: rf,
		out:      o,
	}
}

var fileDirectory = "/pkg/shared/document/photo"

func (i *PhotoInteractor) PhotoCreate(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		tracing.LogError(sp, status.Error(http.StatusBadRequest, "request nil"))
		return nil, status.Error(http.StatusBadRequest, "request nil")
	}

	reqdata := in.(*CreatePhotoRequest)

	reqFileLocal := &entity.SaveFileLocalRequest{
		FileBase64: reqdata.FileBase64,
		FileName:   reqdata.FileName,
		TypeFile:   reqdata.TypeFile,
		PathFile:   fileDirectory,
	}

	resSaveFile, err := i.repoFile.SaveFile(sp, reqFileLocal)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	outSaveFile := resSaveFile.(*entity.SaveFileLocalResponse)

	reqFile := &entity.CreatePhotoRequest{
		FileName:    outSaveFile.FileName,
		TypeFile:    outSaveFile.TypeFile,
		PathFile:    outSaveFile.PathFile,
		Description: reqdata.Description,
	}

	resSave, err := i.repo.PhotoCreate(sp, reqFile)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, resSave)
	return i.out.PhotoCreateResponse(resSave)
}

func (i *PhotoInteractor) PhotoGet(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		return nil, status.Error(http.StatusBadRequest, "request nil")
	}

	reqdata := in.(*PhotoListRequest)

	reqPhoto := &entity.GetPhotoRequest{
		CurPage:   reqdata.CurPage,
		Limit:     reqdata.Limit,
		NamePhoto: reqdata.NamePhoto,
	}
	resdb, err := i.repo.PhotoGet(sp, reqPhoto)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	outPhoto := resdb.(*entity.GetPhotoResponse)

	var photo []Photo
	for _, item := range outPhoto.Data {
		reqFile := &entity.GetFileLocalRequest{
			FileName: item.FileName,
			TypeFile: item.TypeFile,
			PathFile: item.PathFile,
		}
		resFile, errFile := i.repoFile.GetFile(sp, reqFile)
		if errFile != nil {
			tracing.LogError(sp, errFile)
			return nil, errFile
		}
		outFile := resFile.(*entity.GetFileLocalResponse)
		photo = append(photo, Photo{
			Id:          item.Id,
			FileBase64:  outFile.FileBase64,
			FileName:    outFile.FileName,
			TypeFile:    outFile.TypeFile,
			Description: item.Description,
		})
	}
	res := &PhotoListResponse{
		Data:             photo,
		PaginationHelper: outPhoto.PaginationHelper,
	}

	tracing.LogResponse(sp, res)
	return i.out.PhotoGetResponse(res)
}

func (i *PhotoInteractor) PhotoUpdate(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		return nil, status.Error(http.StatusBadRequest, "request nil")
	}

	reqdata := in.(*UpdatePhotoRequest)

	reqDbPhoto := &entity.GetPhotoByIdRequest{Id: reqdata.Id}
	tracing.LogObject(sp, "ReqPhotoGetById", reqDbPhoto)
	resDbPhoto, errDb := i.repo.PhotoGetById(sp, reqDbPhoto)
	if errDb != nil {
		tracing.LogError(sp, errDb)
		return nil, errDb
	}

	outDbPhoto := resDbPhoto.(*entity.GetPhotoByIdResponse)
	reqFileDelete := &entity.DeleteFileLocalRequest{
		FileName: outDbPhoto.FileName,
		TypeFile: outDbPhoto.TypeFile,
		PathFile: outDbPhoto.PathFile,
	}
	tracing.LogObject(sp, "ReqRemoveFile", reqFileDelete)
	_, errDelte := i.repoFile.RemoveFile(sp, reqFileDelete)
	if errDelte != nil {
		tracing.LogError(sp, errDelte)
		return nil, errDelte
	}
	reqFileSave := &entity.SaveFileLocalRequest{
		FileBase64: reqdata.FileBase64,
		FileName:   reqdata.FileName,
		TypeFile:   reqdata.TypeFile,
		PathFile:   fileDirectory,
	}
	tracing.LogObject(sp, "ReqSaveFile", reqFileSave)
	resFileSave, errFileSave := i.repoFile.SaveFile(sp, reqFileSave)
	if errFileSave != nil {
		tracing.LogError(sp, errFileSave)
		return nil, errFileSave
	}
	outFileSave := resFileSave.(*entity.SaveFileLocalResponse)

	reqDbUpdate := &entity.UpdatePhotoRequest{
		Id:          reqdata.Id,
		FileName:    outFileSave.FileName,
		TypeFile:    outFileSave.TypeFile,
		PathFile:    outFileSave.PathFile,
		Description: reqdata.Description,
	}

	tracing.LogObject(sp, "ReqPhotoUpdate", reqDbUpdate)
	resDbUpdate, errDpUpdate := i.repo.PhotoUpdate(sp, reqDbUpdate)
	if errDpUpdate != nil {
		tracing.LogError(sp, errDpUpdate)
		return nil, errDpUpdate
	}

	tracing.LogResponse(sp, resDbUpdate)
	return i.out.PhotoUpdateResponse(resDbUpdate)
}

func (i *PhotoInteractor) PhotoGetById(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		return nil, status.Error(http.StatusBadRequest, "request nil")
	}

	reqdata := in.(*PhotoGetByIdRequest)

	reqDbPhoto := &entity.GetPhotoByIdRequest{Id: reqdata.Id}
	tracing.LogObject(sp, "ReqPhotoGetById", reqDbPhoto)
	resDbPhoto, errDb := i.repo.PhotoGetById(sp, reqDbPhoto)
	if errDb != nil {
		tracing.LogError(sp, errDb)
		return nil, errDb
	}

	outDbPhoto := resDbPhoto.(*entity.GetPhotoByIdResponse)

	reqFile := &entity.GetFileLocalRequest{
		FileName: outDbPhoto.FileName,
		TypeFile: outDbPhoto.TypeFile,
		PathFile: outDbPhoto.PathFile,
	}
	resFile, errFile := i.repoFile.GetFile(sp, reqFile)
	if errFile != nil {
		tracing.LogError(sp, errFile)
		return nil, errFile
	}
	outFile := resFile.(*entity.GetFileLocalResponse)

	res := &PhotoGetByIdResponse{
		Id:          outDbPhoto.Id,
		FileName:    outFile.FileName,
		TypeFile:    outFile.TypeFile,
		FileBase64:  outFile.FileBase64,
		Description: outDbPhoto.Description,
	}

	tracing.LogResponse(sp, res)
	return i.out.PhotoGetByIdResponse(res)
}
func (i *PhotoInteractor) PhotoDeleteById(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	if in == nil {
		return nil, status.Error(http.StatusBadRequest, "request nil")
	}

	reqdata := in.(*DeletePhotoRequest)

	reqDbPhoto := &entity.GetPhotoByIdRequest{Id: reqdata.Id}
	tracing.LogObject(sp, "ReqPhotoGetById", reqDbPhoto)
	resDbPhoto, errDb := i.repo.PhotoGetById(sp, reqDbPhoto)
	if errDb != nil {
		tracing.LogError(sp, errDb)
		return nil, errDb
	}

	outDbPhoto := resDbPhoto.(*entity.GetPhotoByIdResponse)

	reqFileDelete := &entity.DeleteFileLocalRequest{
		FileName: outDbPhoto.FileName,
		TypeFile: outDbPhoto.TypeFile,
		PathFile: outDbPhoto.PathFile,
	}
	tracing.LogObject(sp, "ReqRemoveFile", reqFileDelete)
	_, errDelte := i.repoFile.RemoveFile(sp, reqFileDelete)
	if errDelte != nil {
		tracing.LogError(sp, errDelte)
		return nil, errDelte
	}

	reqDbDelete := &entity.DeletePhotoRequest{Id: reqdata.Id}
	resDb, errDbDel := i.repo.PhotoDelete(sp, reqDbDelete)
	if errDbDel != nil {
		tracing.LogError(sp, errDbDel)
		return nil, errDbDel
	}
	outFile := resDb.(*entity.DeletePhotoResponse)

	tracing.LogResponse(sp, outFile)
	return i.out.PhotoDeleteByIdResponse(outFile)
}
