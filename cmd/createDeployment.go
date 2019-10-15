package cmd

import (
	"fmt"
	"github.com/sap-commerce-tools/ccv2ctl/portal"
	"github.com/spf13/cobra"
	"strings"
)

func commaSeparatedKeys(input map[string]struct{}) string {

	keys := make([]string, len(input))

	i := 0
	for k := range input {
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

var build, environment, deploymentMode, migrationMode string

var allowedDeploymentModes = commaSeparatedKeys(portal.AllowedDeploymentModes)
var allowedMigrationModes = commaSeparatedKeys(portal.AllowedMigrationModes)

func createDeployment(cmd *cobra.Command, args []string) error {

	_, ok := portal.AllowedDeploymentModes[deploymentMode]
	if !ok {
		return fmt.Errorf("Deployment mode \"%s\" not recognized. Allowed values: %s\n", deploymentMode, allowedDeploymentModes)
	}
	_, ok = portal.AllowedMigrationModes[migrationMode]
	if !ok {
		return fmt.Errorf("Migration mode \"%s\" not recognized. Allowed values: %s\n", migrationMode, allowedMigrationModes)
	}
	client := createClient()
	defer client.SaveCookieJar()
	r := client.CreateDeployment(environment, migrationMode, deploymentMode, build)

	prettyJson(r, cmd.OutOrStdout())
	return nil
}

func init() {
	createCmd.AddCommand(createDeploymentCmd)

	createDeploymentCmd.Flags().StringVarP(&build, "build", "b", "", "Build to deploy")
	createDeploymentCmd.MarkFlagRequired("build")
	createDeploymentCmd.Flags().StringVarP(&environment, "environment", "e", "", "Target environment")
	createDeploymentCmd.MarkFlagRequired("environment")
	createDeploymentCmd.Flags().StringVarP(&deploymentMode, "deploymentmode", "d", "ROLLING_UPDATE", "Deployment mode: "+allowedDeploymentModes)
	createDeploymentCmd.Flags().StringVarP(&migrationMode, "migrationmode", "m", "NONE", "Migration mode: "+allowedMigrationModes)
}
