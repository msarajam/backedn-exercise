package responses

type Meta struct {
	Status string   `json:"status"`
	Errors []string `json:"errors,omitempty"`
}
