package webber

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type BasicRequest struct {
	httpRequest *http.Request
	pathParams  map[string]string
}

func NewRequest(r *http.Request) *BasicRequest {
	return &BasicRequest{
		httpRequest: r,
		pathParams:  mux.Vars(r),
	}
}

func (r *BasicRequest) Header(key string) string {
	return r.httpRequest.Header.Get(key)
}

func (r *BasicRequest) PathParam(key string) (string, bool) {
	v, ok := r.pathParams[key]
	return v, ok
}

func (r *BasicRequest) JSON(target interface{}) error {
	return json.NewDecoder(r.httpRequest.Body).Decode(target)
}
