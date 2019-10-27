package api

import (
	"golang.org/x/time/rate"
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

var(
	// 1 operation per each connection
	limiter = rate.NewLimiter(1, 1)
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
	if limiter.Allow() == false {
		return NewResponse(http.StatusTooManyRequests, core.MediaTypeJSON).Writer
	}
	id, ok := req.PathParam(pathParamID)
	if !ok {
		panic("did not receive required path parameter " + pathParamID)
	}
	app, err := c.collection.Fetch(id)
	if err != nil {
		if err == storage.ErrNotFound {
			return NewResponse(http.StatusNotFound, core.MediaTypeJSON).Writer
		}
		log.Println(err)
		return NewResponse(http.StatusInternalServerError, core.MediaTypeJSON).Writer
	}

	/*TODO*/
	//return NewResponse(http.StatusOK, core.MediaTypeYAML).Data(responseKeyApp, app).Writer
		return NewResponse(http.StatusOK, core.MediaTypeJSON).Data(responseKeyApp, app).Writer
}

// Search gets multiple app from storage and returns it
func (c appsController) Search(req core.Request) core.ResponseWriter {
	if limiter.Allow() == false {
		return NewResponse(http.StatusTooManyRequests, core.MediaTypeJSON).Writer
	}
	id, ok := req.PathParam(pathParamID)
	if !ok {
		panic("did not receive required path parameter " + pathParamID)
	}
	app, err := c.collection.Search(id)
	if err != nil {
		if err == storage.ErrNotFound {
			return NewResponse(http.StatusNotFound, core.MediaTypeJSON).Writer
		}
		return NewResponse(http.StatusInternalServerError, core.MediaTypeJSON).Writer
	}
	return NewResponse(http.StatusOK, core.MediaTypeJSON).Data(responseKeyApp, app).Writer
}

// Create adds an app to storage and returns it with its unique identifier
func (c appsController) Create(req core.Request) core.ResponseWriter {
	if limiter.Allow() == false {
		return NewResponse(http.StatusTooManyRequests, core.MediaTypeJSON).Writer
	}
	app := models.App{}
	if err := req.Initialize(&app); err != nil {
		return NewResponse(http.StatusBadRequest, core.MediaTypeJSON).Data(responseKeyErrors, errorBadBody).Writer
	}
	ok, messages, err := app.Validate(c.validate)
	if err != nil {
		return NewResponse(http.StatusInternalServerError, core.MediaTypeJSON).Writer
	}

	if !ok {
		return NewResponse(http.StatusBadRequest, core.MediaTypeJSON).Data(responseKeyErrors, messages).Writer
	}

	app.ID = c.collection.Insert(app)
	return NewResponse(http.StatusCreated, core.MediaTypeJSON).Data(responseKeyApp, app).Writer
}
