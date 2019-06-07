package api

import (
	. "github.com/onsi/gomega"
	"github.com/upbound/backend-exercise/pkg/webber/core"
	"net/http"
	"testing"
)

func TestNewResponse(t *testing.T) {
	type args struct {
		statusCode int
		mediaType  string
	}
	tests := []struct {
		name string
		args args
		want *Response
	}{
		{
			name: "OKJSON",
			args: args{
				statusCode: http.StatusOK,
				mediaType:  core.MediaTypeJSON,
			},
			want: &Response{statusCode: 200,
				mediaType: "application/json",
				data:      map[string]interface{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGomegaWithT(t)

			got := NewResponse(tt.args.statusCode, tt.args.mediaType)
			g.Expect(got).To(Equal(tt.want))
		})
	}
}
