package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var build, environment, mode string
var allowedModes = map[string]struct{}{
	"initialize": {},
	"update":     {},
	"none":       {},
}

func allowedOut() string {
	keys := make([]string, len(allowedModes))

	i := 0
	for k := range allowedModes {
		keys[i] = k
		i++
	}
	return strings.Join(keys, ", ")
}

var createDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "Deploy a build to an environment",
	Long: `Deploy a build to an environment.

Example:

  ccv2ctl create deployment --build 20181001.1 --environment d1

`,
	RunE: createDeployment,
}

func createDeployment(cmd *cobra.Command, args []string) error {

	_, ok := allowedModes[mode]
	if !ok {
		return fmt.Errorf("Mode \"%s\" not recognized. Allowed values: %s\n", mode, allowedOut())
	}
	client := createClient()
	defer client.SaveCookieJar()
	r := client.CreateDeployment(environment, mode, build)

	prettyJson(r, cmd.OutOrStdout())
	return nil
}

func init() {
	createCmd.AddCommand(createDeploymentCmd)

	createDeploymentCmd.Flags().StringVarP(&build, "build", "b", "", "Build to deploy")
	createDeploymentCmd.MarkFlagRequired("build")
	createDeploymentCmd.Flags().StringVarP(&environment, "environment", "e", "", "Target environment")
	createDeploymentCmd.MarkFlagRequired("environment")
	createDeploymentCmd.Flags().StringVarP(&mode, "mode", "m", "none", "Deployment mode: "+allowedOut())

}
