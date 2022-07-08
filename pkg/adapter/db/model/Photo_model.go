package model

type PhotoModel struct {
	BaseModels
	BaseCUModels
	FileName    string `gorm:"column:file_name;unique"`
	TypeFile    string `gorm:"column:type_file"`
	PathFile    string `gorm:"column:path_file;unique"`
	Description string `gorm:"column:description"`
}
