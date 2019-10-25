package main

import (
	"github.com/upbound/backend-exercise/pkg/api"
	"github.com/upbound/backend-exercise/pkg/storage"
	"gopkg.in/go-playground/validator.v8"
	"log"
)

func main() {
	listenAddress := "localhost:8080"
	collection := storage.NewCollection()
	validate := validator.New(&validator.Config{
		TagName:      "validate",
		FieldNameTag: "json",
	})

	log.Printf("Serving on %s", listenAddress)
	if err := api.Serve(listenAddress, collection, validate); err != nil {
		panic(err)
	}
}
