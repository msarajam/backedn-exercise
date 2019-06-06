package core

import (
	"net/http"
)

const (
	HeaderContentType         = "Content-Type"
	HeaderXContentTypeOptions = "X-Content-Type-Options"
	NoSniff                   = "nosniff"
	MediaTypeJSON             = "application/json"
	MethodGet                 = "GET"
	MethodPost                = "POST"
)

type Request interface {
	PathParam(key string) (string, bool)
	JSON(target interface{}) error
}

type ResponseWriter func(w http.ResponseWriter)

type Handler func(r Request) ResponseWriter

type Server interface {
	GET(path string, handler Handler)
	POST(path string, handler Handler)
	Serve() error
}

type Response interface {
	SetErrors([]string)
}
