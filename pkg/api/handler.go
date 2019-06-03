package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/upbound/backend-exercise/pkg/models"
	"github.com/upbound/backend-exercise/pkg/storage"
	"gopkg.in/go-playground/validator.v8"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	internalServerError = "internal server error"
	notFound            = "not found"
	yamlParseError      = "invalid request body YAML"
	idKey               = "id"
)

type handler struct {
	collection *storage.Collection
	validate   *validator.Validate
}

// newHandler creates a new handler.
func newHandler(c *storage.Collection, v *validator.Validate) *handler {
	return &handler{
		collection: c,
		validate:   v,
	}
}

// create adds a new application to storage.
func (h *handler) create(c *gin.Context) {
	app := models.App{}

	if err := h.parseYaml(c.Request.Body, &app); err != nil {
		c.YAML(http.StatusBadRequest, errorMessage(yamlParseError))
		return
	}

	valid, validationMessages, err := h.validateApp(app)
	if err != nil {
		c.YAML(http.StatusInternalServerError, errorMessage(internalServerError))
		return
	}
	if !valid {
		c.YAML(http.StatusBadRequest, errorsList(validationMessages))
		return
	}

	id := h.collection.Insert(app)
	app.ID = id

	c.YAML(http.StatusCreated, app)
}

// list returns all of the applications from storage.
func (h *handler) list(c *gin.Context) {
	c.YAML(http.StatusOK, h.collection.All())
}

func (h *handler) fetch(c *gin.Context) {
	id := c.Param(idKey)
	app, err := h.collection.Fetch(id)
	if err != nil {
		if err == storage.NotFound {
			c.YAML(http.StatusNotFound, errorMessage(notFound))
			return
		}
		c.YAML(http.StatusInternalServerError, errorMessage(internalServerError))
		return
	}
	c.YAML(http.StatusOK, app)
}

// update updates the properties of an application that already exists in storage.
func (h *handler) update(c *gin.Context) {
	app := models.App{}

	if err := h.parseYaml(c.Request.Body, &app); err != nil {
		c.YAML(http.StatusBadRequest, errorMessage(yamlParseError))
		return
	}

	valid, validationMessages, err := h.validateApp(app)
	if err != nil {
		c.YAML(http.StatusInternalServerError, errorMessage(internalServerError))
		return
	}
	if !valid {
		c.YAML(http.StatusBadRequest, errorsList(validationMessages))
		return
	}

	app.ID = c.Param(idKey)
	if err := h.collection.Update(app); err != nil {
		if err == storage.NotFound {
			c.YAML(http.StatusNotFound, errorMessage(notFound))
			return
		}
		c.YAML(http.StatusInternalServerError, errorMessage(internalServerError))
		return
	}

	c.YAML(http.StatusOK, app)
}


// remove deletes an existing application from storage.
func (h *handler) remove(c *gin.Context) {
	id := c.Param(idKey)
	if err := h.collection.Remove(id); err != nil {
		if err == storage.NotFound {
			c.YAML(http.StatusNotFound, errorMessage(notFound))
			return
		}
		c.YAML(http.StatusInternalServerError, errorMessage(internalServerError))
	}
}

// parseYaml reads raw YAML input and parses it into the given target variable,
func (h *handler) parseYaml(closer io.ReadCloser, target interface{}) error {
	raw, err := ioutil.ReadAll(closer)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(raw, target)
}

// validationErrorsMessage takes an error and parses it into a list of helpful validation messages.
func (h *handler) validationErrorsMessage(e error) (map[string][]string, error) {
	fields, ok := e.(validator.ValidationErrors)
	if !ok {
		return nil, errors.New("unable to convert error to type ValidationErrors")
	}

	list := make([]string, len(fields))
	i := 0
	for _, f := range fields {
		list[i] = fmt.Sprintf("Field '%s' failed validation '%s'", f.Field, f.Tag)
		i++
	}
	return map[string][]string{"errors": list}, nil
}

// validateApp validates an application and returns three things: bool (is valid), []string (validation error messages), error
func (h *handler) validateApp(a models.App) (bool, []string, error) {
	if err := h.validate.Struct(a); err != nil {
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
