package api

import (
	"github.com/upbound/backend-exercise/pkg/server"
	"github.com/upbound/backend-exercise/pkg/server/core"
	"github.com/upbound/backend-exercise/pkg/storage"
	"gopkg.in/go-playground/validator.v8"
)

func Serve(listenAddress string, c *storage.Collection, v *validator.Validate) error {
	s := server.NewServer(listenAddress, core.MediaTypeJSON)
	apps := newAppsController(c, v)

	s.POST("/apps", apps.Create)
	s.GET("/apps/{id}", apps.Fetch)

	return s.Serve()
}
