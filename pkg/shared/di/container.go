package di

import (
	"github.com/sarulabs/di"
	"news/pkg/adapter/db/connection"
	"news/pkg/adapter/repository/db"
	"news/pkg/adapter/repository/file"
	dto "news/pkg/infrastructure/restful/service/dto"
	"news/pkg/usecase/authentication"
	"news/pkg/usecase/kategori"
	"news/pkg/usecase/photo"
	"news/pkg/usecase/user"
)

// MetaData :
type MetaData struct {
	ClientName string
	ClientIP   string
	UUID       string
	ActivityID string
}

// Container :
type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add([]di.Def{
		{Name: "kategori", Build: kategoriUsecase},
		{Name: "user", Build: userUsecase},
		{Name: "photo", Build: photoUsecase},
		{Name: "authentication", Build: authenticationUsecase},
	}...)
	return &Container{
		ctn: builder.Build(),
	}

}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func kategoriUsecase(_ di.Container) (interface{}, error) {
	repo := db.NewKategoryDatahandler(connection.NewsDB)
	out := &dto.KategoriBuilder{}
	return kategori.NewKatergoriInteractor(repo, out), nil
}

func userUsecase(_ di.Container) (interface{}, error) {
	repo := db.NewUserDataHandler(connection.NewsDB)
	out := &dto.UserBuilder{}
	return user.NewUserInteractor(repo, out), nil
}

func photoUsecase(_ di.Container) (interface{}, error) {
	repo := db.NewPhotoDataHanlder(connection.NewsDB)
	repoFIle := file.NewFileLocalDataHandler()
	out := &dto.PhotoBuilder{}
	return photo.NewPhotoInteractor(repo, repoFIle, out), nil
}

func authenticationUsecase(_ di.Container) (interface{}, error) {
	repo := db.NewAuthenticationDataHandler(connection.NewsDB)
	out := &dto.AuthenticationBuilder{}
	return authentication.NewAuthenticationInteractor(repo, out), nil
}
