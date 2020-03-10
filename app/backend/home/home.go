package home

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/qortenant/backend"
	"go-tenancy/config/application"
)

// New new home app
func New(config *Config) *App {
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
}

// ConfigureApplication configure application
func (App) ConfigureApplication(application *application.Application) {
	AdminParty := application.IrisApplication.AdminParty
	Iris := application.IrisApplication.Iris
	backend.Register(Iris)

	AdminParty.HandleDir("/static", "app/backend/home/views/static")
	Iris.RegisterView(iris.HTML("./app/backend/home/views", ".html"))

	AdminParty.Get("/", func(ctx iris.Context) {
		if err := ctx.View("index.html"); err != nil {
			color.Red(fmt.Sprintf("Home Index View error: %v\n", err))
		}
	})

}
