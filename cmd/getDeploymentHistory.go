package cmd

import (
	"github.com/spf13/cobra"
)

var deploymentHistoryCmd = &cobra.Command{
	Use:   "deploymenthistory [environment]",
	Short: "Get details about deployments.",
	Long: `Get details about deployments on an environment.
The command expects the environment code as the single argument.

Example:

  ccv2ctl get deploymenthistory d1

`,
	Args:    cobra.ExactArgs(1),
	Run:     getAndPrintDeploymentHistory,
}

func getAndPrintDeploymentHistory(cmd *cobra.Command, args []string) {
	client := createClient()
	defer client.SaveCookieJar()

	running := client.GetDeployments(args[0])
	prettyJson(running, cmd.OutOrStdout())
}

func init() {
	getCmd.AddCommand(deploymentHistoryCmd)
}
