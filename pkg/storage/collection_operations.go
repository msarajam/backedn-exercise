package storage

import "github.com/upbound/backend-exercise/pkg/models"

//go:generate mockgen -destination ./mocks/mock_collection_operations.go -package mocks github.com/upbound/backend-exercise/pkg/storage CollectionOperations
type CollectionOperations interface {
	Insert(a models.App) string
	Fetch(id string) (models.App, error)
}
