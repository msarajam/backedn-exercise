package webber

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
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
	err := json.NewDecoder(r.httpRequest.Body).Decode(target)
	fmt.Println("in JSON err :", err)
	return err
}

func (r *BasicRequest) YAML(target interface{}) error {
	err := yaml.NewDecoder(r.httpRequest.Body).Decode(target)
	/*TODO*/
	fmt.Println("in YAML err :", err)
	return nil
}
