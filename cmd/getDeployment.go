 package cmd

import (
	"github.com/spf13/cobra"
)

var deploymentCmd = &cobra.Command{
	Use:   "deployment [environment]",
	Short: "Get details about running deployments(s).",
	Long: `Get details about running deployments(s) on an environment.
The command expects the environment code as the single argument.

Example:

  ccv2ctl get deployment d1

`,
	Args:    cobra.ExactArgs(1),
	Run:     getAndPrintDeployments,
	Aliases: []string{"deployments"},
}

func getAndPrintDeployments(cmd *cobra.Command, args []string) {
	client := createClient()
	defer client.SaveCookieJar()

	running := client.GetRunningDeployments(args[0])
	prettyJson(running, cmd.OutOrStdout())
}

func init() {
	getCmd.AddCommand(deploymentCmd)
}
