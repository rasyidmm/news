package jwtGen

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/status"
	"net/http"
	"news/pkg/shared/tracing"
	"time"
)

type JwtCustomClaims struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	JenisUser string `json:"jenis_user"`
	UUID      string `json:"uuid"`
	jwt.StandardClaims
}
type TokenRequest struct {
	Username  string
	Email     string
	JenisUser string
}

type TokenResponse struct {
	Token   string
	Expired string
}

type JwtClaimResponse struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	JenisUser string `json:"jenis_user"`
	UUID      string `json:"uuid"`
}

var SecretKey = []byte("hala")

func CreateJwtToken(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "CreateJwtToken")
	defer sp.Finish()
	tracing.LogRequest(sp, in)
	reqdata := in.(*TokenRequest)

	Expired := time.Now().Add(5 * time.Hour)

	claims := JwtCustomClaims{
		JenisUser: reqdata.JenisUser,
		Username:  reqdata.Username,
		Email:     reqdata.Email,
		UUID:      uuid.NewString(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: Expired.Unix(),
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(SecretKey)
	if err != nil {
		tracing.LogError(sp, status.Error(http.StatusBadRequest, err.Error()))
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}
	res := &TokenResponse{
		Token:   t,
		Expired: Expired.Format("2006-01-02 15:04:05"),
	}
	return res, nil
}

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	Claims:     &JwtCustomClaims{},
	SigningKey: SecretKey,
})

func GetClientMetadata(c echo.Context) (JwtClaimResponse, error) {
	var res JwtClaimResponse
	user := c.Get("user").(*jwt.Token)
	claim := user.Claims.(*JwtCustomClaims)

	if claim.Username == "" {
		return res, c.JSON(http.StatusUnauthorized, "")
	} else if claim.Email == "" {
		return res, c.JSON(http.StatusUnauthorized, "")
	} else if claim.UUID == "" {
		return res, c.JSON(http.StatusUnauthorized, "")
	} else if claim.JenisUser == "" {
		return res, c.JSON(http.StatusUnauthorized, "")
	}
	res = JwtClaimResponse{
		Username:  claim.Username,
		Email:     claim.Email,
		JenisUser: claim.JenisUser,
		UUID:      claim.UUID,
	}
	return res, nil
}
