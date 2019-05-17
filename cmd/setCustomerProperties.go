package cmd

import (
	"github.com/spf13/cobra"
)

var setPropAspect, setPropFilename string

var setCustomerPropertiesCmd = &cobra.Command{
	Use:   "customerproperties [environment]",
	Short: "Set customer properties.",
	Long: `Set customer properties for specific aspect of an environment.
The command expects the environment code as the single argument.

Example:

  ccv2ctl set customerproperties d1 -a hcs_aspect -f properties

`,
	Args:    cobra.ExactArgs(1),
	Run:     setCustomerProperties,
}

func setCustomerProperties(cmd *cobra.Command, args []string) {
	client := createClient()
	defer client.SaveCookieJar()

	customerProperties := client.SetCustomerProperties(args[0], setPropAspect, setPropFilename)
	prettyJson(customerProperties, cmd.OutOrStdout())
}

func init() {
	setCmd.AddCommand(setCustomerPropertiesCmd)

	setCustomerPropertiesCmd.Flags().StringVarP(&setPropAspect, "aspect", "a", "", "Name of the aspect (required)")
	setCustomerPropertiesCmd.MarkFlagRequired("aspect")
	setCustomerPropertiesCmd.Flags().StringVarP(&setPropFilename, "propertyfile", "f", "-", "Name of the propertyfile")
}
