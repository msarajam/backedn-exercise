package api

import (
	"github.com/gin-gonic/gin"
	"github.com/upbound/backend-exercise/pkg/storage"
	"gopkg.in/go-playground/validator.v8"
)

// Serve starts an HTTP API server and listens for incoming requests.
func Serve(c *storage.Collection, v *validator.Validate) error {
	r := gin.Default()
	h := newHandler(c, v)

	r.GET("/apps", h.list)
	r.GET("/apps/:id", h.fetch)
	r.POST("/apps", h.create)
	r.PUT("/apps/:id", h.update)
	r.DELETE("/apps/:id", h.remove)

	return r.Run() // serve on 0.0.0.0:8080
}
