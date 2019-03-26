package cmd

import (
	"github.com/spf13/cobra"
)

var initialpasswordCmd = &cobra.Command{
	Use:   "initialpassword [environment]",
	Short: "Get initial passwords.",
	Long: `Get initial passwors of an environment.
The command expects the environment code as the single argument.

Example:

  ccv2ctl get initialpassword d1

`,
	Args:    cobra.ExactArgs(1),
	Run:     getInitialPasswords,
	Aliases: []string{"initialpasswords"},
}

func getInitialPasswords(cmd *cobra.Command, args []string) {
	client := createClient()
	defer client.SaveCookieJar()

	passwords := client.GetInitialPasswords(args[0])
	prettyJson(passwords, cmd.OutOrStdout())
}

func init() {
	getCmd.AddCommand(initialpasswordCmd)
}
