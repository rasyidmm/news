package file

import (
	"bytes"
	"encoding/base64"
	"github.com/nfnt/resize"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/status"
	"image"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"news/pkg/domain/entity"
	"news/pkg/shared/tracing"
	"os"
	"strings"
)

type FileLocalDataHandler struct {
}

func NewFileLocalDataHandler() *FileLocalDataHandler {
	return &FileLocalDataHandler{}
}

func (f *FileLocalDataHandler) SaveFile(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "UserCreate")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.SaveFileLocalRequest)

	dir, err := os.Getwd()
	if err != nil {
		tracing.LogError(sp, status.Error(http.StatusBadRequest, err.Error()))
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}

	err = os.MkdirAll(dir+reqdata.PathFile, os.ModePerm)
	if err != nil {
		tracing.LogError(sp, status.Error(http.StatusBadRequest, err.Error()))
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}

	if err != nil {
		tracing.LogError(sp, status.Error(http.StatusBadRequest, "create diractory gagal"))
		return nil, status.Error(http.StatusBadRequest, "create diractory gagal")
	}

	var opts jpeg.Options

	unbased, err := base64.StdEncoding.DecodeString(reqdata.FileBase64)
	if err != nil {
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}

	img, _, err := image.Decode(bytes.NewReader(unbased))
	if err != nil {
		tracing.LogError(sp, status.Error(http.StatusBadRequest, err.Error()))
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}

	realPath := reqdata.PathFile + "/" + reqdata.FileName + "." + reqdata.TypeFile
	out, _ := os.Create(dir + realPath)
	defer out.Close()
	m := resize.Resize(480, 320, img, resize.Lanczos3)

	if strings.ToUpper(reqdata.TypeFile) == "JPG" {
		opts.Quality = 100
		err = jpeg.Encode(out, m, &opts)
		if err != nil {
			tracing.LogError(sp, status.Error(http.StatusBadRequest, err.Error()))
			return nil, status.Error(http.StatusBadRequest, err.Error())
		}
	}

	res := &entity.SaveFileLocalResponse{
		FileName: reqdata.FileName,
		TypeFile: reqdata.TypeFile,
		PathFile: realPath,
	}
	tracing.LogResponse(sp, res)
	return res, nil
}

func (f *FileLocalDataHandler) GetFile(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetFile")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.GetFileLocalRequest)
	dir, err := os.Getwd()
	if err != nil {
		tracing.LogError(sp, status.Error(http.StatusBadRequest, err.Error()))
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}

	bytes, err := ioutil.ReadFile(dir + reqdata.PathFile)
	if err != nil {
		tracing.LogError(sp, status.Error(http.StatusBadRequest, err.Error()))
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}

	var base64Encoding string

	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += base64.StdEncoding.EncodeToString(bytes)

	res := &entity.GetFileLocalResponse{
		FileBase64: base64Encoding,
		FileName:   reqdata.FileName,
		TypeFile:   reqdata.TypeFile,
		PathFile:   reqdata.PathFile,
	}
	tracing.LogResponse(sp, res)
	return res, nil
}
func (f *FileLocalDataHandler) UpdateFile(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "UpdateFile")
	defer sp.Finish()
	tracing.LogRequest(sp, in)
	return nil, nil
}

func (f *FileLocalDataHandler) RemoveFile(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "RemoveFile")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*entity.DeleteFileLocalRequest)

	dir, err := os.Getwd()
	if err != nil {
		tracing.LogError(sp, status.Error(http.StatusBadRequest, err.Error()))
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}

	e := os.Remove(dir + reqdata.PathFile)
	if !os.IsNotExist(e) {
		tracing.LogError(sp, status.Error(http.StatusBadRequest, e.Error()))
		return nil, status.Error(http.StatusBadRequest, e.Error())
	}

	res := entity.DeleteFileLocalResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
