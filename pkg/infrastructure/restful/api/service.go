package api

import (
	"fmt"
	Apis "github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"log"
	cfg "news/internal/config"
	authentication_service "news/pkg/infrastructure/restful/service/authentication"
	comment_service "news/pkg/infrastructure/restful/service/comment"
	kategor_service "news/pkg/infrastructure/restful/service/kategori"
	news_service "news/pkg/infrastructure/restful/service/news"
	photo_service "news/pkg/infrastructure/restful/service/photo"
	user_service "news/pkg/infrastructure/restful/service/user"
	"news/pkg/infrastructure/router"
	container "news/pkg/shared/di"
	"news/pkg/usecase/authentication"
	"news/pkg/usecase/comment"
	"news/pkg/usecase/news"
	"news/pkg/usecase/photo"

	//_ "news/pkg/shared/document/swagger"
	"news/pkg/shared/tracing"
	"news/pkg/usecase/kategori"
	"news/pkg/usecase/user"
)

// // @title Echo Swagger Example API
//// @version 1.0
//// @description This is a sample server server.
//// @termsOfService http://swagger.io/terms/
//
//// @contact.name API Support
//// @contact.url http://www.swagger.io/support
//// @contact.email support@swagger.io
//
//// @license.name Apache 2.0
//// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
//// @host localhost:3000
//// @BasePath /
//// @schemes http
func RunServer() {
	log.Println("Starting Restfull Server")

	config := cfg.GetConfig()

	fmt.Println(config)
	e := Apis.New()
	ctn := container.NewContainer()

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
	router.NewPhotoRouter(e, photo_service.NewPhotoService(ctn.Resolve("photo").(*photo.PhotoInteractor)))
	router.NewAuthenticationRouter(e, authentication_service.NewAuthenticationService(ctn.Resolve("authentication").(*authentication.AuthenticationInteractor)))
	router.NewNewsRouter(e, news_service.NewNewsService(ctn.Resolve("news").(*news.NewsInteractor)))
	router.NewCommentRouter(e, comment_service.NewCommentService(ctn.Resolve("comment").(*comment.CommentInteractor)))
}
