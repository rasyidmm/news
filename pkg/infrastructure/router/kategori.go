package router

import (
	service "Home/news/pkg/infrastructure/restful/service/kategori"
)

type KategoriRouter struct {
	serv service.KategoriService
}
