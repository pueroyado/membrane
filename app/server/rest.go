package server

import (
	"context"
	"demo/controllers"
	_ "demo/docs"
	"demo/repositories"
	"demo/utils"
	"fmt"
	"github.com/jmoiron/sqlx"
	httpSwagger "github.com/swaggo/http-swagger"
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
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	userRepo := repositories.NewUserRepo(s.dbMysql)
	userHandler := controllers.NewHandlerUser(userRepo)
	r.HandleFunc("/user/reg", userHandler.Reg()).Methods(http.MethodPost)
	r.HandleFunc("/user/auth", userHandler.Auth()).Methods(http.MethodPost)

	secure := r.PathPrefix("/api").Subrouter()
	secure.Use(utils.JwtVerify)

	productRepo := repositories.NewProductRepo(s.dbMysql)
	handlerProduct := controllers.NewHandlerProduct(productRepo)
	secure.HandleFunc("/product", handlerProduct.List()).Methods(http.MethodGet)
	secure.HandleFunc("/product/{id:[0-9]+}", handlerProduct.Detail()).Methods(http.MethodGet)

	return r
}

func (s *APIServer) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is home page!")
	}
}
