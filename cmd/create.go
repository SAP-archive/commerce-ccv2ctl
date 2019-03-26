package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create CCv2 resources",
}

func init() {
	rootCmd.AddCommand(createCmd)
}
