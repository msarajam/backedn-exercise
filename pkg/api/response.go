package api

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"github.com/upbound/backend-exercise/pkg/webber/core"
	"log"
	"net/http"
)

// Response is the structure of response msg
type Response struct {
	statusCode int
	mediaType  string
	data       map[string]interface{}
}

// NewResponse is for creating a new response
func NewResponse(statusCode int, mediaType string) *Response {
	return &Response{
		statusCode: statusCode,
		mediaType:  mediaType,
		data:       map[string]interface{}{},
	}
}

//Data is being use in the NewResponse for structuring the data
func (r *Response) Data(key string, value interface{}) *Response {
	r.data[key] = value
	return r
}

// Writer is for writing the response
func (r *Response) Writer(w http.ResponseWriter) {
	// Write the header first (important!)
	r.writeHeader(w)
	// If there is no data we are done here
	if len(r.data) == 0 {
		return
	}
	// Write the body
	r.writeBody(w, r.marshal())
}

func (r *Response) writeHeader(w http.ResponseWriter) {
	w.Header().Set(core.HeaderContentType, r.mediaType)
	w.Header().Set(core.HeaderXContentTypeOptions, core.NoSniff)
	w.WriteHeader(r.statusCode)
}

func (r *Response) writeBody(w http.ResponseWriter, body []byte) {
	if _, err := w.Write(body); err != nil {
		log.Println(err)
	}
}

// marshal the data to the appropriate media type
func (r *Response) marshal() []byte {
	var body []byte
	switch r.mediaType {
	case core.MediaTypeJSON:
		body = r.marshalJSON()
	case core.MediaTypeYAML:
		body = r.marshalYaml()
	default:
		panic("unsupported media type: " + r.mediaType)
	}
	return body
}

// marshal the data to JSON
func (r *Response) marshalJSON() []byte {
	parsed, err := json.Marshal(r.data)
	if err != nil {
		panic(err)
	}
	return parsed
}

// marshal the data to YAML
func (r *Response) marshalYaml() []byte {
	parsed, err := yaml.Marshal(r.data)
	if err != nil {
		panic(err)
	}
	return parsed
}
