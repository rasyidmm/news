package model

type UserModel struct {
	BaseModels
	BaseCUModels
	Username       string `gorm:"column:username;unique"`
	FirstName      string `gorm:"column:first_name"`
	LastName       string `gorm:"column:last_name"`
	Twitter        string `gorm:"column:twitter"`
	Facebook       string `gorm:"column:facebook"`
	Instagram      string `gorm:"column:instagram"`
	Biography      string `gorm:"column:biography"`
	Email          string `gorm:"column:email;unique"`
	NomerHandphone string `gorm:"column:nomer_handphone;unique"`
	Password       string `gorm:"column:password"`
	JenisUser      string `gorm:"column:jenis_user"`
}
