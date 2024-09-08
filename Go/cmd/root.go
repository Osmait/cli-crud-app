package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hello",
	Short: "Una aplicación CLI de ejemplo",
	Long:  `Esta es una aplicación de línea de comandos básica de ejemplo usando Cobra en Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hola, esta es una CLI de ejemplo en Go!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
