package entity

type SaveFileLocalRequest struct {
	FileBase64 string
	FileName   string
	TypeFile   string
	PathFile   string
}
type SaveFileLocalResponse struct {
	FileName string
	TypeFile string
	PathFile string
}

type GetFileLocalRequest struct {
	FileName string
	TypeFile string
	PathFile string
}
type GetFileLocalResponse struct {
	FileBase64 string
	FileName   string
	TypeFile   string
	PathFile   string
}

type DeleteFileLocalRequest struct {
	FileName string
	TypeFile string
	PathFile string
}
type DeleteFileLocalResponse struct {
	StatusCode string
	StatusDesc string
}
