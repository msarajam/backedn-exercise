package webber

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

func jsonCheck(b []byte, target interface{}) error {
	return json.Unmarshal([]byte(b), target)
}

func yamlCheck(b []byte, target interface{}) error {
	return yaml.Unmarshal([]byte(b), target)
}

func (r *BasicRequest) Initialize(target interface{}) error {
	b, err := ioutil.ReadAll(r.httpRequest.Body)
	if err != nil {
		return err
	}

	//check the json format
	if err = jsonCheck(b, target); err == nil {
		return nil
	}

	//check the yaml format
	if err = yamlCheck(b, target); err == nil {
		return nil
	}

	return errors.New("the requested data is not valid")
}
