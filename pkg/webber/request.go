package webber

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/upbound/backend-exercise/pkg/webber/core"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

// BasicRequest is the request struct
type BasicRequest struct {
	httpRequest *http.Request
	pathParams  map[string]string
}

// NewRequest is for creating new request
func NewRequest(r *http.Request) *BasicRequest {
	return &BasicRequest{
		httpRequest: r,
		pathParams:  mux.Vars(r),
	}
}

// Header to get the header from http request
func (r *BasicRequest) Header(key string) string {
	return r.httpRequest.Header.Get(key)
}

// PathParam is getting the request msg
func (r *BasicRequest) PathParam(key string) (string, string, bool) {
	msgType := ""
	switch r.httpRequest.RequestURI[:len("/apps/yaml/")] {
	case "/apps/yaml/":
		msgType = core.MediaTypeYAML
	default:
		msgType = core.MediaTypeJSON
	}
	v, ok := r.pathParams[key]
	return msgType, v, ok
}

func jsonCheck(b []byte, target interface{}) error {
	return json.Unmarshal([]byte(b), target)
}

func yamlCheck(b []byte, target interface{}) error {
	return yaml.Unmarshal([]byte(b), target)
}

// Initialize is for checking the json and yaml,
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
