package cmd

import (
	"github.com/spf13/cobra"
)

var aspect string

var customerPropertiesCmd = &cobra.Command{
	Use:   "customerproperties [environment]",
	Short: "Get customer properties.",
	Long: `Get customer properties for specific aspect of an environment.
The command expects the environment code as the single argument.

Example:

  ccv2ctl get customerproperties d1 -a hcs_aspect

`,
	Args:    cobra.ExactArgs(1),
	Run:     getAndPrintCustomerProperties,
}

func getAndPrintCustomerProperties(cmd *cobra.Command, args []string) {
	client := createClient()
	defer client.SaveCookieJar()

	customerProperties := client.GetCustomerProperties(args[0], aspect)
	prettyJson(customerProperties, cmd.OutOrStdout())
}

func init() {
	getCmd.AddCommand(customerPropertiesCmd)

	customerPropertiesCmd.Flags().StringVarP(&aspect, "aspect", "a", "", "Name of the aspect (required)")
	customerPropertiesCmd.MarkFlagRequired("aspect")
}
