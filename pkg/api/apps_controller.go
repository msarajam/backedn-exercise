package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/upbound/backend-exercise/pkg/models"
	"github.com/upbound/backend-exercise/pkg/storage"
	"github.com/upbound/backend-exercise/pkg/webber/core"
	"gopkg.in/go-playground/validator.v8"
)

const (
	pathParamID       = "id"
	responseKeyApp    = "app"
	responseKeyErrors = "errors"
	errorBadBody      = "invalid request body"
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
	fmt.Println("in GET", 1)
	id, ok := req.PathParam(pathParamID)
	if !ok {
		panic("did not receive required path parameter " + pathParamID)
	}

	fmt.Println("in GET", 2)
	app, err := c.collection.Fetch(id)
	if err != nil {

		fmt.Println("in GET", 3, err)
		if err == storage.ErrNotFound {
			return NewResponse(http.StatusNotFound, core.MediaTypeJSON).Writer
		}
		log.Println(err)
		return NewResponse(http.StatusInternalServerError, core.MediaTypeJSON).Writer
	}

	fmt.Println("in GET", "last")
	return NewResponse(http.StatusOK, core.MediaTypeJSON).Data(responseKeyApp, app).Writer
}

// Create adds an app to storage and returns it with its unique identifier
func (c appsController) Create(req core.Request) core.ResponseWriter {
	app := models.App{}
	jsonType := true
	if err := req.JSON(&app); err != nil {
		jsonType = false
	}
	if !jsonType {
		fmt.Println(jsonType)
		if err := req.YAML(&app); err != nil {
			fmt.Println("Got an error in yaml")
			return NewResponse(http.StatusBadRequest, core.MediaTypeJSON).Data(responseKeyErrors, errorBadBody).Writer
		}
	}
	fmt.Println("HERERERERE")
	ok, messages, err := app.Validate(c.validate)
	fmt.Println("HERE", ok, "\n", messages, "\n", err)
	if err != nil {
		log.Println(err)
		return NewResponse(http.StatusInternalServerError, core.MediaTypeJSON).Writer
	}

	if !ok {
		return NewResponse(http.StatusBadRequest, core.MediaTypeJSON).Data(responseKeyErrors, messages).Writer
	}

	app.ID = c.collection.Insert(app)

	return NewResponse(http.StatusCreated, core.MediaTypeJSON).Data(responseKeyApp, app).Writer
}
