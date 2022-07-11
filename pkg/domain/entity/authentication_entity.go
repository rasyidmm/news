package entity

type CheckUsernamePasswordRequest struct {
	Username string
	Password string
}

type CheckUsernamePasswordResponse struct {
	Username       string
	FirstName      string
	LastName       string
	Twitter        string
	Facebook       string
	Instagram      string
	Biography      string
	Email          string
	NomerHandphone string
	Password       string
	JenisUser      string
	UserId         string
}

type LoginHistorySaveRequest struct {
	Username    string
	UserId      string
	Email       string
	JenisUser   string
	ExpiredTime string
	IpAddress   string
}

type LoginHistorySaveResponse struct {
	StatusCode string
	StatusDesc string
}
