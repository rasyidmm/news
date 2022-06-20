package model

type KategoriModel struct {
	BaseModels
	BaseCUModels
	KategoriName string `gorm:"column:kategori_name;unique"`
	KategoriCode string `gorm:"column:kategori_code;unique"`
}
