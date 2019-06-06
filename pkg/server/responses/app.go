package responses

import "github.com/upbound/backend-exercise/pkg/models"

type App struct {
	Response
	App models.App `json:"app"`
}

func NewApp(a models.App) *App {
	return &App{App: a}
}
