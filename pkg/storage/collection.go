package storage

import (
	uuid "github.com/satori/go.uuid"
	"github.com/upbound/backend-exercise/pkg/models"
	"strings"
)

// Collection is the App struct
type Collection struct {
	data map[string]models.App
}

// NewCollection is for creating new App struct
func NewCollection() *Collection {
	return &Collection{
		data: map[string]models.App{},
	}
}

// Insert is for saving data in App struct
func (c *Collection) Insert(a models.App) string {
	a.ID = uuid.NewV4().String()
	c.data[a.ID] = a
	return a.ID
}

// Fetch is for get the data that has been saved in App struct
func (c *Collection) Fetch(id string) (models.App, error) {
	if a, ok := c.data[id]; ok {
		return a, nil
	}
	return models.App{}, ErrNotFound
}

// Search is for searching the given value (http request) in Collection data
func (c *Collection) Search(searchParameter string) ([]models.App, error) {
	modelApp := []models.App{}
	sID := strings.ToLower(searchParameter[:strings.Index(searchParameter, "=")])
	sParameter := strings.ToLower(searchParameter[len(sID)+1:])
	for _, v := range c.data {
		switch sID {
		case "id":
			if strings.Contains(strings.ToLower(v.ID), sParameter) {
				modelApp = append(modelApp, v)
			}
		case "title":
			if strings.Contains(strings.ToLower(v.Title), sParameter) {
				modelApp = append(modelApp, v)
			}
		case "maintainers":
			for _, vm := range v.Maintainers {
				if strings.Contains(strings.ToLower(vm.Email), sParameter) {
					modelApp = append(modelApp, v)
				} else if strings.Contains(strings.ToLower(vm.Name), sParameter) {
					modelApp = append(modelApp, v)
				}
			}
		case "company":
			if strings.Contains(strings.ToLower(v.Company), sParameter) {
				modelApp = append(modelApp, v)
			}
		case "website":
			if strings.Contains(strings.ToLower(v.Title), sParameter) {
				modelApp = append(modelApp, v)
			}
		case "source":
			if strings.Contains(strings.ToLower(v.Source), sParameter) {
				modelApp = append(modelApp, v)
			}
		case "license":
			if strings.Contains(strings.ToLower(v.License), sParameter) {
				modelApp = append(modelApp, v)
			}
		case "description":
			if strings.Contains(strings.ToLower(v.Description), sParameter) {
				modelApp = append(modelApp, v)
			}
		}
	}
	if len(modelApp) == 0 {
		return modelApp, ErrNotFound
	}
	return modelApp, nil
}

func checkMaintainer() {

}
