package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build [code] OR build --all",
	Short: "Get details of a cloud build",
	Long: `Get details of a cloud build.

Either specifiy the code of the build to get a single build or use the --all flag to get all builds

Examples:

  ccv2ctl get build 20180930.3
  # get the data of a single build

  ccv2ctl get build --all
  # get meta data of all builds of the subscription

`,
	RunE:    getAndPrintBuild,
	Args:    cobra.MaximumNArgs(1),
	Aliases: []string{"builds"},
}

func getAndPrintBuild(cmd *cobra.Command, args []string) error {
	getAll, _ := cmd.Flags().GetBool("all")

	if !getAll && len(args) == 0 {
		return fmt.Errorf("Neither --all flag nor <code> specified. Exiting\n")
	}

	client := createClient()
	defer client.SaveCookieJar()

	var output interface{}
	if getAll {
		output = client.GetAllBuilds()
	} else {
		output = client.GetBuild(args[0])
	}
	prettyJson(output, cmd.OutOrStdout())
	return nil
}

func init() {
	getCmd.AddCommand(buildCmd)

	buildCmd.Flags().BoolP("all", "a", false, "Get all builds")
}
