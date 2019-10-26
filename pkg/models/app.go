package models

import (
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v8"
	"strings"
)

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

// Validate the app - returns valid, validation messages, error
func (a App) Validate(v *validator.Validate) (bool, []string, error) {
	if err := v.Struct(a); err != nil {
		fields, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, nil, errors.New("unable to convert error to type ValidationErrors")
		}

		list := make([]string, len(fields))
		i := 0
		for _, f := range fields {
			names := strings.Split(f.NameNamespace, ".")[1:]
			name := strings.Join(names, ".")
			list[i] = fmt.Sprintf("Field '%s' failed validation rule '%s'", name, f.Tag)
			i++
		}
		return false, list, nil
	}
	return true, nil, nil
}
