package server

import (
	"github.com/gorilla/mux"
	"github.com/upbound/backend-exercise/pkg/core"
	"github.com/upbound/backend-exercise/pkg/server/middleware"
	"net/http"
	"time"
)

type Server struct {
	router     *mux.Router
	httpServer *http.Server
	mediaType  string
}

func NewServer(listenAddress, mediaType string) *Server {
	r := mux.NewRouter()
	r.Use(middleware.ResponseContentType(mediaType))
	r.MethodNotAllowedHandler = methodNotAllowedHandler(mediaType)

	return &Server{
		router: r,
		httpServer: &http.Server{
			Addr:         listenAddress,
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  10 * time.Second,
		},
		mediaType: mediaType,
	}
}

func (s *Server) GET(path string, h core.Handler) {
	s.register(path, h, core.MethodGet)
}

func (s *Server) POST(path string, h core.Handler) {
	s.register(path, h, core.MethodPost)
}

func (s *Server) Serve() error {
	s.httpServer.Handler = s.router
	return s.httpServer.ListenAndServe()
}

func (s *Server) register(path string, h core.Handler, method string) {
	s.router.HandleFunc(path, wrap(h)).Methods(method)
}