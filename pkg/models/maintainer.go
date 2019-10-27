package models

// Maintainer is a struct to use in the App json struct
type Maintainer struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}
