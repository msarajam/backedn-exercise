package apps

import (
	"errors"
	"fmt"
	"github.com/upbound/backend-exercise/pkg/models"
	"github.com/upbound/backend-exercise/pkg/server"
	"github.com/upbound/backend-exercise/pkg/server/core"
	"github.com/upbound/backend-exercise/pkg/server/responses"
	"github.com/upbound/backend-exercise/pkg/storage"
	"gopkg.in/go-playground/validator.v8"
	"log"
	"strings"
)

var (
	Collection *storage.Collection
	Validate   *validator.Validate
)

// Fetch gets a single app from storage and returns it
func Fetch(req core.Request) core.ResponseWriter {
	id, ok := req.PathParam("id")
	if !ok {
		err := fmt.Errorf("did not receive required path parameter '%s'", "id")
		return server.InternalServerError(err)
	}

	app, err := Collection.Fetch(id)
	if err != nil {
		if err == storage.NotFound {
			return server.NotFound()
		}
		return server.InternalServerError(err)
	}

	return server.OK(responses.NewApp(app))
}

// Create adds an app to storage and returns it with its unique identifier
func Create(req core.Request) core.ResponseWriter {
	app := models.App{}
	if err := req.JSON(&app); err != nil {
		log.Println(err)
		return server.BadRequest([]string{"invalid request body"})
	}

	ok, messages, err := validateApp(app)
	if err != nil {
		return server.InternalServerError(err)
	}

	if !ok {
		return server.BadRequest(messages)
	}

	app.ID = Collection.Insert(app)

	return server.Created(responses.NewApp(app))
}

// validateApp take an app and returns: bool (is valid), []string (validation error messages), error
func validateApp(a models.App) (bool, []string, error) {
	if err := Validate.Struct(a); err != nil {
		fields, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, nil, errors.New("unable to convert error to type ValidationErrors")
		}

		list := make([]string, len(fields))
		i := 0
		for _, f := range fields {
			names := strings.Split(f.NameNamespace, ".")[1:]
			name := strings.Join(names, ".")
			list[i] = fmt.Sprintf("Field '%s' failed validation rule '%s'", name, f.Tag)
			i++
		}
		return false, list, nil
	}
	return true, nil, nil
}
