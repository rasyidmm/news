package api

import (
	cfg "Home/news/internal/config"
	container "Home/news/pkg/shared/di"
	"fmt"
	Apis "github.com/labstack/echo/v4"
	"log"
)

func RunServer() {
	log.Println("Starting Restfull Server")

	config := cfg.GetConfig()

	fmt.Println(config)
	e := Apis.New()
	ctn := container.NewContainer()
	Apply(e, ctn)
	svcPort := config.Server.Rest.Port

	e.Logger.Fatal(e.Start(":" + svcPort))
}

func Apply(e *Apis.Echo, ctn *container.Container) {

}
