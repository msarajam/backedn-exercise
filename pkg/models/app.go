package models

type App struct {
	ID          string       `json:"id"`
	Title       string       `json:"title" validate:"required"`
	Version     string       `json:"version" validate:"required"`
	Maintainers []Maintainer `json:"maintainers" validate:"required,dive"`
	Company     string       `json:"company" validate:"required"`
	Website     string       `json:"website" validate:"required"`
	Source      string       `json:"source" validate:"required"`
	License     string       `json:"license" validate:"required"`
	Description string       `json:"description" validate:"required"`
}
