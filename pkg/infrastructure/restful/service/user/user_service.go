package user

import (
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"news/pkg/shared/tracing"
	usecase "news/pkg/usecase/user"
)

type UserCreateRequest struct {
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Twitter        string `json:"twitter"`
	Facebook       string `json:"facebook"`
	Instagram      string `json:"instagram"`
	Biography      string `json:"biography"`
	Email          string `json:"email"`
	NomerHandphone string `json:"nomer_handphone"`
	Password       string `json:"password"`
	JenisUser      string `json:"jenis_user"`
}

type UserListRequest struct {
	CurPage int `json:"cur_page"`
	Limit   int `json:"limit"`
}

type User struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Twitter        string `json:"twitter"`
	Facebook       string `json:"facebook"`
	Instagram      string `json:"instagram"`
	Biography      string `json:"biography"`
	Email          string `json:"email"`
	NomerHandphone string `json:"nomer_handphone"`
	Password       string `json:"password"`
	JenisUser      string `json:"jenis_user"`
}

type UserGetByIdRequest struct {
	UserId string `json:"user_id"`
}

type UserUpdateRequest struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Twitter        string `json:"twitter"`
	Facebook       string `json:"facebook"`
	Instagram      string `json:"instagram"`
	Biography      string `json:"biography"`
	Email          string `json:"email"`
	NomerHandphone string `json:"nomer_handphone"`
	Password       string `json:"password"`
	JenisUser      string `json:"jenis_user"`
}

type UserDeleteRequest struct {
	UserId string `json:"user_id"`
}

type UserService struct {
	uc usecase.UserInputPort
}

func NewUserService(u usecase.UserInputPort) *UserService {
	return &UserService{
		uc: u,
	}
}

func (s *UserService) UserCreate(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "UserCreate")
	defer sp.Finish()

	reqdata := new(UserCreateRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.UserCreateRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.UserCreate(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)

}

func (s *UserService) UserList(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "UserCreate")
	defer sp.Finish()

	reqdata := new(UserListRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.UserListRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.UserList(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (s *UserService) UserGetById(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "UserCreate")
	defer sp.Finish()

	id := c.Param("id")
	reqdata := &UserGetByIdRequest{
		UserId: id,
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.UserGetByIdRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.UserGetById(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (s *UserService) UserUpdate(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "UserUpdate")
	defer sp.Finish()

	id := c.Param("id")
	reqdata := new(UserUpdateRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	request := &usecase.UserUpdateRequest{
		Id:             id,
		Username:       reqdata.Username,
		FirstName:      reqdata.FirstName,
		LastName:       reqdata.LastName,
		Twitter:        reqdata.Twitter,
		Facebook:       reqdata.Facebook,
		Instagram:      reqdata.Instagram,
		Biography:      reqdata.Biography,
		Email:          reqdata.Email,
		NomerHandphone: reqdata.NomerHandphone,
		Password:       reqdata.Password,
		JenisUser:      reqdata.JenisUser,
	}

	tracing.LogRequest(sp, request)
	res, err := s.uc.UserUpdate(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (s *UserService) UserDelete(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "UserDelete")
	defer sp.Finish()

	id := c.Param("id")
	reqdata := &UserDeleteRequest{
		UserId: id,
	}
	tracing.LogRequest(sp, reqdata)

	var request *usecase.UserDeleteRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.UserDelete(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}
