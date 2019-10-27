package webber

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/upbound/backend-exercise/pkg/webber/core"
)

// Server is the structure of the Server
type Server struct {
	router     *mux.Router
	httpServer *http.Server
	mediaType  string
}

// NewServer is for creating new server
func NewServer(listenAddress, mediaType string) *Server {
	r := mux.NewRouter()
	r.NotFoundHandler = notFoundHandler(mediaType)
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

// GET is the http get request to send data
func (s *Server) GET(path string, h core.Handler) {
	s.register(path, h, core.MethodGet)
}

// POST is http post request to receive data
func (s *Server) POST(path string, h core.Handler) {
	s.register(path, h, core.MethodPost)
}

// Serve to start the listener
func (s *Server) Serve() error {
	s.httpServer.Handler = s.router
	return s.httpServer.ListenAndServe()
}

func (s *Server) register(path string, h core.Handler, method string) {
	s.router.HandleFunc(path, wrap(h)).Methods(method)
}
