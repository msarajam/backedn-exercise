package main

import (
	"github.com/upbound/backend-exercise/pkg/api"
	"github.com/upbound/backend-exercise/pkg/storage"
	"gopkg.in/go-playground/validator.v8"
)

func main() {
	collection := storage.NewCollection()

	config := &validator.Config{
		TagName:      "validate",
		FieldNameTag: "yaml",
	}

	validate := validator.New(config)

	if err := api.Serve(collection, validate); err != nil {
		panic(err)
	}
}
