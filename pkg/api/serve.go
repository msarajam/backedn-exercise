package api

import (
	"github.com/upbound/backend-exercise/pkg/storage"
	"github.com/upbound/backend-exercise/pkg/webber"
	"github.com/upbound/backend-exercise/pkg/webber/core"
	"gopkg.in/go-playground/validator.v8"
)

func Serve(listenAddress string, c *storage.Collection, v *validator.Validate) error {
	s := webber.NewServer(listenAddress, core.MediaTypeJSON)
	apps := newAppsController(c, v)

	s.POST("/apps", apps.Create)
	s.GET("/apps/json/{id}", apps.Fetch)
	s.GET("/apps/yaml/{id}", apps.Fetch)
	s.GET("/apps/search/{id}", apps.Search)
	return s.Serve()
}
