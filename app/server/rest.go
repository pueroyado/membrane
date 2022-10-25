package server

import (
	"context"
	"demo/controllers/product"
	"demo/repositories"
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type APIServer struct {
	router     *mux.Router
	serverHttp *http.Server
	dbMysql    *sqlx.DB
}

func Create() *APIServer {
	return &APIServer{
		router:  mux.NewRouter(),
		dbMysql: NewMysqlConnect(),
	}
}

func (s *APIServer) Start() error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: s.Router(),
	}
	s.serverHttp = server

	return server.ListenAndServe()
}

func (s *APIServer) Shutdown(mainCtx context.Context) error {
	ctx, cancel := context.WithTimeout(mainCtx, time.Duration(3)*time.Second)
	defer cancel()

	err := s.serverHttp.Shutdown(ctx)
	return err
}

func (s *APIServer) Router() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", s.handleHome()).Methods(http.MethodGet)

	productRepo := repositories.NewProductRepo(s.dbMysql)
	handlerProduct := product.NewHandlerProduct(productRepo)
	r.HandleFunc("/product", handlerProduct.List()).Methods(http.MethodGet)
	r.HandleFunc("/product", product.Create()).Methods(http.MethodPost)
	r.HandleFunc("/product", product.Delete()).Methods(http.MethodDelete)
	r.HandleFunc("/product", product.Update()).Methods(http.MethodPatch)

	return r
}

func (s *APIServer) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is home page!")
	}
}
