package authentication

import (
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"news/pkg/shared/tracing"
	usecase "news/pkg/usecase/authentication"
)

type LoginRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Ipaddress string `json:"ip_address"`
}

type AuthenticationService struct {
	uc usecase.AuthenticationInputPort
}

func NewAuthenticationService(u usecase.AuthenticationInputPort) *AuthenticationService {
	return &AuthenticationService{
		uc: u,
	}
}

func (s *AuthenticationService) Login(c echo.Context) error {
	sp, _ := tracing.CreateRootSpan(c, "KategoriCreate")
	defer sp.Finish()

	reqdata := new(LoginRequest)
	if err := c.Bind(reqdata); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogRequest(sp, reqdata)

	var request *usecase.LoginRequest
	err := mapstructure.Decode(reqdata, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := s.uc.Login(c, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}
