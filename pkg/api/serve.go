package api

import (
	"fmt"

	"github.com/upbound/backend-exercise/pkg/storage"
	"github.com/upbound/backend-exercise/pkg/webber"
	"github.com/upbound/backend-exercise/pkg/webber/core"
	"gopkg.in/go-playground/validator.v8"
)

func Serve(listenAddress string, c *storage.Collection, v *validator.Validate) error {
	s := webber.NewServer(listenAddress, core.MediaTypeJSON)
	apps := newAppsController(c, v)

	fmt.Println("in Server")
	s.POST("/apps", apps.Create)
	s.GET("/apps/{id}", apps.Fetch)

	fmt.Println("going to Server")
	return s.Serve()
}
