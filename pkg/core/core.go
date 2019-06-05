package core

import (
	"encoding/json"
	"log"
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

func NotFound() ResponseWriter {
	return func(w http.ResponseWriter) {
		w.WriteHeader(http.StatusNotFound)
	}
}

func BadRequest() ResponseWriter {
	return func(w http.ResponseWriter) {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func InternalServerError(err error) ResponseWriter {
	log.Println(err)
	return func(w http.ResponseWriter) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func OK(body interface{}) ResponseWriter {
	parsed, err := json.Marshal(body)
	if err != nil {
		return InternalServerError(err)
	}

	return func(w http.ResponseWriter) {
		w.WriteHeader(http.StatusOK)

		if _, err := w.Write(parsed); err != nil {
			log.Println(err)
		}
	}
}

func Created(body interface{}) ResponseWriter {
	parsed, err := json.Marshal(body)
	if err != nil {
		return InternalServerError(err)
	}

	return func(w http.ResponseWriter) {
		w.WriteHeader(http.StatusCreated)

		if _, err := w.Write(parsed); err != nil {
			log.Println(err)
		}
	}
}
