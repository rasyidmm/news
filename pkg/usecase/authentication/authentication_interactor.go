package authentication

import (
	"github.com/labstack/echo/v4"
	"news/pkg/domain/entity"
	"news/pkg/domain/repository"
	"news/pkg/shared/enum"
	"news/pkg/shared/jwtGen"
	"news/pkg/shared/tracing"
)

type LoginRequest struct {
	Username  string
	Password  string
	Ipaddress string
}
type LoginResponse struct {
	Token     string
	Username  string
	Expired   string
	JenisUser string
}
type AuthenticationInteractor struct {
	repo repository.AuthenticationRepository
	out  AuthenticationOutputPort
}

func NewAuthenticationInteractor(r repository.AuthenticationRepository, o AuthenticationOutputPort) *AuthenticationInteractor {
	return &AuthenticationInteractor{
		repo: r,
		out:  o,
	}
}

func (i *AuthenticationInteractor) Login(ctx echo.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartInteractor))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqdata := in.(*LoginRequest)

	reqCheck := &entity.CheckUsernamePasswordRequest{
		Username: reqdata.Username,
		Password: reqdata.Password,
	}

	tracing.LogObject(sp, "CheckUsernamePassword", reqCheck)
	resCek, errCek := i.repo.CheckUsernamePassword(sp, reqCheck)
	if errCek != nil {
		tracing.LogError(sp, errCek)
		return nil, errCek
	}

	outCek := resCek.(*entity.CheckUsernamePasswordResponse)
	reqJwt := &jwtGen.TokenRequest{
		Username:  outCek.Username,
		Email:     outCek.Email,
		JenisUser: outCek.JenisUser,
	}

	tracing.LogObject(sp, "CreateJwtToken", reqJwt)
	resTok, errTok := jwtGen.CreateJwtToken(sp, reqJwt)
	if errTok != nil {
		tracing.LogError(sp, errTok)
		return nil, errTok
	}

	outTok := resTok.(*jwtGen.TokenResponse)
	reqHistoryLogin := &entity.LoginHistorySaveRequest{
		Username:    reqdata.Username,
		UserId:      outCek.UserId,
		Email:       outCek.Email,
		JenisUser:   outCek.JenisUser,
		ExpiredTime: outTok.Expired,
		IpAddress:   reqdata.Ipaddress,
	}

	tracing.LogObject(sp, "LoginHistorySave", reqHistoryLogin)
	i.repo.LoginHistorySave(sp, reqHistoryLogin)

	res := &LoginResponse{
		Token:     outTok.Token,
		Username:  outCek.Username,
		Expired:   outTok.Expired,
		JenisUser: outCek.JenisUser,
	}

	tracing.LogResponse(sp, res)
	return i.out.LoginResponse(res)
}
