package api

import (
	"fmt"
	"github.com/upbound/backend-exercise/pkg/storage"
	"gopkg.in/go-playground/validator.v8"
)

func Serve(c *storage.Collection, v *validator.Validate) error {
	fmt.Println("herrow\ngoodbye")
	return nil
}
