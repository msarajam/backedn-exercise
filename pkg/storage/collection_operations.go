package storage

import "github.com/upbound/backend-exercise/pkg/models"

// CollectionOperations the interface to use in collection.go
type CollectionOperations interface {
	Insert(a models.App) string
	Fetch(id string) (models.App, error)
}
