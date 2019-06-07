package webber

import (
	"github.com/upbound/backend-exercise/pkg/webber/core"
	"net/http"
)

func wrap(handler core.Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := NewRequest(r)
		resp := handler(req)
		resp(w)
	}
}

func notFoundHandler(mediaType string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(core.HeaderContentType, mediaType)
		w.Header().Set(core.HeaderXContentTypeOptions, core.NoSniff)
		w.WriteHeader(http.StatusNotFound)
	})
}

func methodNotAllowedHandler(mediaType string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(core.HeaderContentType, mediaType)
		w.Header().Set(core.HeaderXContentTypeOptions, core.NoSniff)
		w.WriteHeader(http.StatusMethodNotAllowed)
	})
}
