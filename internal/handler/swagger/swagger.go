package swagger

import (
	"github.com/MuZaZaVr/account-service/internal/handler"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

func HandlerSwagger(router *handler.API) {
	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)

	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./docs/")))
}
