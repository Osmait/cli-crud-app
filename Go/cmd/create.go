package cmd

import (
	"github.com/osmait/cli-crud-app/services"
	"github.com/spf13/cobra"
)

func NewCreateCommant(noteServices *services.NoteServices) *cobra.Command {
	createdCmd := &cobra.Command{
		Use:   "create [nombre]",
		Short: "Saluda a una persona",
		Long:  `Este comando permite saludar a una persona por su nombre.`,
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]

			description := args[1]
			noteServices.Create(name, description)
		},
	}
	return createdCmd
}
