package model

type LoginHistoryModel struct {
	BaseModels
	BaseCUModels
	Username    string `gorm:"column:username"`
	UserId      string `gorm:"column:user_id"`
	Email       string `gorm:"column:email"`
	JenisUser   string `gorm:"column:jenis_user"`
	ExpiredTime string `gorm:"column:expired_time"`
	IpAddress   string `gorm:"column:ip_address"`
}
