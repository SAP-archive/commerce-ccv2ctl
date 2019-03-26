package cmd

import (
	"github.com/spf13/cobra"
)

var name, branch string

var createBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Create and start a new build",
	Long: `Create and start a new build.
The build is automatically triggered, the response is the build meta data. 
The "Code" of the build in the response the unique identifier used for all other build related commands (get build, logs)

The branch name is not verified against your repository when creating the build, so please double check

Example:

  ccv2ctl create build --branch some-branch --name build-name

`,
	Run: createBuild,
}

func createBuild(cmd *cobra.Command, args []string) {
	client := createClient()
	defer client.SaveCookieJar()
	newBuild := client.CreateBuild(name, branch)

	prettyJson(newBuild, cmd.OutOrStdout())
}

func init() {
	createCmd.AddCommand(createBuildCmd)

	createBuildCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the build (required)")
	createBuildCmd.MarkFlagRequired("name")
	createBuildCmd.Flags().StringVarP(&branch, "branch", "b", "", "Branch to build (required)")
	createBuildCmd.MarkFlagRequired("branch")
}
