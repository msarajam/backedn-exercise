package responses

type Response struct {
	Errors []string `json:"errors,omitempty"`
}

func NewResponse() Response {
	return Response{}
}

func (r *Response) SetErrors(e []string) {
	r.Errors = e
}
