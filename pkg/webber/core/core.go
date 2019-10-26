package core

import "net/http"

const (
	HeaderContentType         = "Content-Type"
	HeaderXContentTypeOptions = "X-Content-Type-Options"
	NoSniff                   = "nosniff"
	MediaTypeJSON             = "application/json"
	MethodGet                 = "GET"
	MethodPost                = "POST"
)

type ResponseWriter func(w http.ResponseWriter)

type Handler func(r Request) ResponseWriter

type Request interface {
	PathParam(key string) (string, bool)
	Initialize(target interface{}) error
	Header(key string) string
}
