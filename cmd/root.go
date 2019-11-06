package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gochecks",
	Short: "",
	//SilenceErrors: true,
	SilenceUsage: true,
	// TODO mb PreRunE env var check
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
