package api

import (
	"fmt"
	"github.com/upbound/backend-exercise/pkg/models"
	"github.com/upbound/backend-exercise/pkg/server"
	"github.com/upbound/backend-exercise/pkg/server/core"
	"github.com/upbound/backend-exercise/pkg/server/responses"
	"github.com/upbound/backend-exercise/pkg/storage"
	"gopkg.in/go-playground/validator.v8"
	"log"
)

type appsController struct {
	collection *storage.Collection
	validate   *validator.Validate
}

func newAppsController(c *storage.Collection, v *validator.Validate) appsController {
	return appsController{
		collection: c,
		validate:   v,
	}
}

// Fetch gets a single app from storage and returns it
func (c appsController) Fetch(req core.Request) core.ResponseWriter {
	id, ok := req.PathParam("id")
	if !ok {
		err := fmt.Errorf("did not receive required path parameter '%s'", "id")
		return server.InternalServerError(err)
	}

	app, err := c.collection.Fetch(id)
	if err != nil {
		if err == storage.ErrNotFound {
			return server.NotFound()
		}
		return server.InternalServerError(err)
	}

	return server.OK(responses.NewApp(app))
}

// Create adds an app to storage and returns it with its unique identifier
func (c appsController) Create(req core.Request) core.ResponseWriter {
	app := models.App{}
	if err := req.JSON(&app); err != nil {
		log.Println(err)
		return server.BadRequest([]string{"invalid request body"})
	}

	ok, messages, err := app.Validate(c.validate)
	if err != nil {
		return server.InternalServerError(err)
	}

	if !ok {
		return server.BadRequest(messages)
	}

	app.ID = c.collection.Insert(app)

	return server.Created(responses.NewApp(app))
}
