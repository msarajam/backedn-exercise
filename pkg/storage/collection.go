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
	id := uuid.NewV4().String()
	a.ID = id
	c.data[id] = a
	return id
}

func (c *Collection) Fetch(id string) (models.App, error) {
	a, ok := c.data[id]
	if !ok {
		return models.App{}, NotFound
	}
	return a, nil
}

func (c *Collection) All() []models.App {
	apps := make([]models.App, len(c.data))
	i := 0
	for _, app := range c.data {
		apps[i] = app
		i++
	}
	return apps
}

func (c *Collection) Update(a models.App) error {
	_, err := c.Fetch(a.ID)
	if err != nil {
		return err
	}
	c.data[a.ID] = a
	return nil
}

func (c *Collection) Remove(id string) error {
	_, err := c.Fetch(id)
	if err != nil {
		return err
	}
	delete(c.data, id)
	return nil
}
