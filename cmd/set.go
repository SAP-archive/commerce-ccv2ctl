package cmd

import (
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set values of CCv2 resources",
}

func init() {
	rootCmd.AddCommand(setCmd)
}
