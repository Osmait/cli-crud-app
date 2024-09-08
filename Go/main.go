package main

import (
	"github.com/osmait/cli-crud-app/cmd"
	"github.com/osmait/cli-crud-app/repository"
	"github.com/osmait/cli-crud-app/services"
)

func main() {
	appRepository, _ := repository.NewNoteRepository()
	appServices := services.NewNoteServices(appRepository)
	app := cmd.NewApp(appServices)
	app.Execute()
}
