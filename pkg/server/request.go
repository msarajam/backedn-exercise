package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type request struct {
	body        []byte
	httpRequest *http.Request
	pathParams  map[string]string
}

func NewRequest(r *http.Request) *request {
	return &request{
		httpRequest: r,
		pathParams:  mux.Vars(r),
	}
}

func (r *request) PathParam(key string) (string, bool) {
	v, ok := r.pathParams[key]
	return v, ok
}

func (r *request) Body() ([]byte, error) {
	if r.body == nil {
		raw, err := ioutil.ReadAll(r.httpRequest.Body)
		if err != nil {
			return nil, err
		}
		r.body = raw
	}
	return r.body, nil
}

func (r *request) JSON(target interface{}) error {
	raw, err := r.Body()
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, target)
}
