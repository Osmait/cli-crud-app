package cmd

import "github.com/osmait/cli-crud-app/services"

type App struct {
	appServices *services.NoteServices
}

func NewApp(appServices *services.NoteServices) *App {
	return &App{
		appServices: appServices,
	}
}

func (a *App) Commants() {
	createdCmd := NewCreateCommant(a.appServices)
	rootCmd.AddCommand(createdCmd)
}

func (a *App) Execute() {
	a.Commants()
	Execute()
}
