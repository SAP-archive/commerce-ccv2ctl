package cmd

import (
	"github.com/spf13/cobra"
	"io"
)

var logsCmd = &cobra.Command{
	Use:   "logs [code]",
	Short: "Dump build log to stdout",
	Long: `Download, unzip and dump build log file to stdout
The command expects the build code as single argument

Example:

  ccv2ctl logs 20180930.2

`,
	Run:  dumpBuildLog,
	Args: cobra.ExactArgs(1),
}

func dumpBuildLog(cmd *cobra.Command, args []string) {

	client := createClient()
	defer client.SaveCookieJar()

	reader := client.GetBuildLogReader(args[0])

	io.Copy(cmd.OutOrStdout(), reader)
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
