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
	return models.App{}, NotFound
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
	if _, err := c.Fetch(a.ID); err != nil {
		return err
	}

	c.data[a.ID] = a
	return nil
}

func (c *Collection) Remove(id string) error {
	if _, err := c.Fetch(id); err != nil {
		return err
	}

	delete(c.data, id)
	return nil
}
