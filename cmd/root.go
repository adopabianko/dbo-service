package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dbo-service",
	Short: "DBO Service",
	Run: func(cmd *cobra.Command, args []string) {
		StartHTTPServer()
	},
}

func init() {
	// add command
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal("failed execute command", err)
	}
}
