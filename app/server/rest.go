package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type APIServer struct {
	router     *mux.Router
	serverHttp *http.Server
}

func Create() *APIServer {
	return &APIServer{
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: s.RouterConf(),
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

func (s *APIServer) RouterConf() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", s.handleHome()).Methods(http.MethodGet)
	r.HandleFunc("/info", s.handleInfo()).Methods(http.MethodGet)

	return r
}

func (s *APIServer) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "handleInfo")
	}
}
func (s *APIServer) handleInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		name := r.URL.Query().Get("name")
		fmt.Fprintln(w, "Hello guy!", name)
	}
}
