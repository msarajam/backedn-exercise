package api

import (
	"encoding/json"
	"github.com/upbound/backend-exercise/pkg/webber/core"
	"log"
	"net/http"
)

type Response struct {
	statusCode int
	mediaType  string
	data       map[string]interface{}
}

func NewResponse(statusCode int, mediaType string) *Response {
	return &Response{
		statusCode: statusCode,
		mediaType:  mediaType,
		data:       map[string]interface{}{},
	}
}

func (r *Response) Data(key string, value interface{}) *Response {
	r.data[key] = value
	return r
}

func (r *Response) Writer(w http.ResponseWriter) {
	// If there is no data, just write the header and bail out
	if len(r.data) == 0 {
		w.WriteHeader(r.statusCode)
		return
	}

	var body []byte

	// Marshal the data to the appropriate media type
	switch r.mediaType {
	case core.MediaTypeJSON:
		parsed, err := json.Marshal(r.data)
		if err != nil {
			panic(err)
		}
		body = parsed
	default:
		panic("unsupported media type: " + r.mediaType)
	}

	w.WriteHeader(r.statusCode)
	if _, err := w.Write(body); err != nil {
		log.Println(err)
	}
}
