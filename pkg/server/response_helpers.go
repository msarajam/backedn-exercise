package server

import (
	"encoding/json"
	"github.com/upbound/backend-exercise/pkg/server/core"
	"github.com/upbound/backend-exercise/pkg/server/responses"
	"log"
	"net/http"
)

func NotFound() core.ResponseWriter {
	r := responses.NewResponse()
	r.SetStatus(http.StatusText(http.StatusNotFound))

	return func(w http.ResponseWriter) {
		w.WriteHeader(http.StatusNotFound)
	}
}

func BadRequest(errs []string) core.ResponseWriter {
	r := responses.NewResponse()
	r.SetStatus(http.StatusText(http.StatusBadRequest))
	r.SetErrors(errs)

	parsed, err := json.Marshal(r)
	if err != nil {
		return InternalServerError(err)
	}

	return func(w http.ResponseWriter) {
		w.WriteHeader(http.StatusBadRequest)

		if _, err := w.Write(parsed); err != nil {
			log.Println(err)
		}
	}
}

func InternalServerError(err error) core.ResponseWriter {
	log.Println(err)

	r := responses.NewResponse()
	r.SetStatus(http.StatusText(http.StatusBadRequest))

	return func(w http.ResponseWriter) {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func OK(r core.Response) core.ResponseWriter {
	r.SetStatus(http.StatusText(http.StatusOK))

	parsed, err := json.Marshal(r)
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

func Created(r core.Response) core.ResponseWriter {
	r.SetStatus(http.StatusText(http.StatusCreated))

	parsed, err := json.Marshal(r)
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
