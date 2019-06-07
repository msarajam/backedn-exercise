package webber

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type BasicRequest struct {
	body        []byte
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

func (r *BasicRequest) Body() ([]byte, error) {
	if r.body == nil {
		raw, err := ioutil.ReadAll(r.httpRequest.Body)
		if err != nil {
			return nil, err
		}
		r.body = raw
	}
	return r.body, nil
}

func (r *BasicRequest) JSON(target interface{}) error {
	raw, err := r.Body()
	if err != nil {
		return err
	}
	return json.Unmarshal(raw, target)
}
