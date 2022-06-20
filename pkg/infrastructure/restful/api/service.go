package api

import (
	"fmt"
	Apis "github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
	_ "news/docs"
	cfg "news/internal/config"
	kategor_service "news/pkg/infrastructure/restful/service/kategori"
	user_service "news/pkg/infrastructure/restful/service/user"
	"news/pkg/infrastructure/router"
	container "news/pkg/shared/di"
	"news/pkg/shared/tracing"
	"news/pkg/usecase/kategori"
	"news/pkg/usecase/user"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used

// @securitydefinitions.oauth2.application  OAuth2Application
// @tokenUrl                                https://example.com/oauth/token
// @scope.write                             Grants write access
// @scope.admin                             Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit  OAuth2Implicit
// @authorizationUrl                     https://example.com/oauth/authorize
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.password  OAuth2Password
// @tokenUrl                             https://example.com/oauth/token
// @scope.read                           Grants read access
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode  OAuth2AccessCode
// @tokenUrl                               https://example.com/oauth/token
// @authorizationUrl                       https://example.com/oauth/authorize
// @scope.admin                            Grants read and write access to administrative information

func RunServer() {
	log.Println("Starting Restfull Server")

	config := cfg.GetConfig()

	fmt.Println(config)
	e := Apis.New()
	ctn := container.NewContainer()
	//c := jaegertracing.New(e, nil)
	//defer c.Close()

	tracer, closer := tracing.Init(e, "news-me", nil)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
	Apply(e, ctn)
	svcPort := config.Server.Rest.Port
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":" + svcPort))
}

func Apply(e *Apis.Echo, ctn *container.Container) {
	router.NewKategoriRouter(e, kategor_service.NewKategoriService(ctn.Resolve("kategori").(*kategori.KatergoriInteractor)))
	router.NewUserRouter(e, user_service.NewUserService(ctn.Resolve("user").(*user.UserInteractor)))
}
