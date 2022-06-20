package api

import (
	"fmt"
	Apis "github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
	cfg "news/internal/config"
	kategor_service "news/pkg/infrastructure/restful/service/kategori"
	user_service "news/pkg/infrastructure/restful/service/user"
	"news/pkg/infrastructure/router"
	container "news/pkg/shared/di"
	"news/pkg/shared/tracing"
	"news/pkg/usecase/kategori"
	"news/pkg/usecase/user"
)

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
