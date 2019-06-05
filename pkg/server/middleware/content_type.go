package middleware

import (
	"github.com/upbound/backend-exercise/pkg/core"
	"net/http"
)

func ResponseContentType(mediaType string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(core.HeaderContentType, mediaType)
			w.Header().Set(core.HeaderXContentTypeOptions, core.NoSniff)
			next.ServeHTTP(w, r)
		})
	}
}
