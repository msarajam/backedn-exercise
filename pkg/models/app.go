package models

type App struct {
	ID          string       `yaml:"id"`
	Title       string       `yaml:"title" validate:"required"`
	Version     string       `yaml:"version" validate:"required"`
	Maintainers []Maintainer `yaml:"maintainers" validate:"required,dive"`
	Company     string       `yaml:"company" validate:"required"`
	Website     string       `yaml:"website" validate:"required"`
	Source      string       `yaml:"source" validate:"required"`
	License     string       `yaml:"license" validate:"required"`
	Description string       `yaml:"description" validate:"required"`
}
