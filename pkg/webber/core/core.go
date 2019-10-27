package core

import "net/http"

// Constants to use in the project
const (
	HeaderContentType         = "Content-Type"
	HeaderXContentTypeOptions = "X-Content-Type-Options"
	NoSniff                   = "nosniff"
	MediaTypeJSON             = "application/json"
	MediaTypeYAML             = "application/x-yaml"
	MethodGet                 = "GET"
	MethodPost                = "POST"
)

// ResponseWriter is going to be use for resounding the request
type ResponseWriter func(w http.ResponseWriter)

// Handler is going to be used in the server
type Handler func(r Request) ResponseWriter

// Request is the struct that server uses to receive the request
type Request interface {
	PathParam(key string) (string, string, bool)
	Initialize(target interface{}) error
	Header(key string) string
}
