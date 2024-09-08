package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	greetCmd := &cobra.Command{
		Use:   "greet [nombre]",
		Short: "Saluda a una persona",
		Long:  `Este comando permite saludar a una persona por su nombre.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			fmt.Printf("Hola, %s!\n", name)
		},
	}

	rootCmd.AddCommand(greetCmd)
}
