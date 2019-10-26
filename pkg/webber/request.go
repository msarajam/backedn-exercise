package webber

import (
	"encoding/json"
	"fmt"
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
	fmt.Println("JSON  :", string(b))
	if err := jsonCheck(b, target); err != nil {
		fmt.Println("Not JSON")
		//return err
	}
	fmt.Println("YAML  :", string(b))
	if err := yamlCheck(b, target); err != nil {
		fmt.Println("Not YAML")
		return err
	}
	return nil
}
