package storage

import (
	"github.com/satori/go.uuid"
	"github.com/upbound/backend-exercise/pkg/models"
)

type Collection struct {
	data map[string]models.App
}

func NewCollection() *Collection {
	return &Collection{
		data: map[string]models.App{},
	}
}

func (c *Collection) Insert(a models.App) string {
	a.ID = uuid.NewV4().String()
	c.data[a.ID] = a
	return a.ID
}

func (c *Collection) Fetch(id string) (models.App, error) {
	if a, ok := c.data[id]; ok {
		return a, nil
	}
	return models.App{}, ErrNotFound
}
