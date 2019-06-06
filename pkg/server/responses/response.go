package responses

type Response struct {
	Meta Meta `json:"meta"`
}

func NewResponse() Response {
	return Response{}
}

func (r *Response) SetStatus(s string) {
	r.Meta.Status = s
}

func (r *Response) SetErrors(e []string) {
	r.Meta.Errors = e
}
