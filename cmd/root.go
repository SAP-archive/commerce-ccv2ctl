package cmd

import (
	"fmt"
	"github.com/sap-commerce-tools/ccv2ctl/portal"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "ccv2ctl",
	Short: "Unofficial CLI for CCv2",
	Long: `Unofficial CLI for CCv2

Check the output of the "help"" command for details and examples of the various commands.

! WARNING: there is (nearly) no validation of the input data              !
!          Make sure you use the correct values for all "create" commands !

The configuration file (default location: $HOME/.ccv2ctl.yaml) allows you to configure:

  certfile: /path/to/certfile.pem
  # Path to PEM-encoded SAP Passport client certificate

  keyfile: /path/to/keyfile.pem
  # Path to PEM-encoded key of SAP Passport client certificate

  subscription: c0deba5ec0deba5ec0deba5ec0deba5e
  # (optional) Default subscription-ID to use for all commands. You can find the ID in the URL of the cloud portal.
  # https://portal.commerce.ondemand.com/subscription/c0deba5ec0deba5ec0deba5ec0deba5e/...
  #                                                   ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

  cookiejar: /path/to/jar
  # (optional) path to HTTP cookie jar
  # Default value: $HOME/.ccv2jar

You can also use environment variables for the above configuration keys, prefixed with "CCV2_", e.g.:
  export CCV2_SUBSCRIPTION=asdfasdf
  ccv2ctl get build 20180930.2

To extract certfile and keyfile out of PKCS#12 encoded store, you can use following OpenSSL commands:
(you will be prompted for the keystore password)

  openssl pkcs12 -in /path/to/store.pfx -nokeys -nodes | openssl x509 -out certfile.pem
  openssl pkcs12 -in /path/to/store.pfx -nocerts -nodes |  openssl rsa -out keyfile.pem

`,
}

func Execute() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
			os.Exit(1)
		}
	}()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func certAndKey() (certPEMBlock, keyPEMBlock []byte) {

	certFile := viper.GetString("certfile")
	certPEMBlock, err := ioutil.ReadFile(certFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not load certificate from '%s'.\n", certFile)
		fmt.Fprintf(os.Stderr, `Check --certfile or "certfile: ..." in $HOME/.ccv2ctl.yaml\n"`)

		os.Exit(1)
	}

	keyFile := viper.GetString("keyfile")
	keyPEMBlock, err = ioutil.ReadFile(keyFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not load key from '%s'.\n", keyFile)
		fmt.Fprint(os.Stderr, `Check --keyfile or "keyfile: ..." in $HOME/.ccv2ctl.yaml\n`)
		os.Exit(1)
	}

	return certPEMBlock, keyPEMBlock
}

func getSubscription() (s string) {
	s = viper.GetString("subscription")
	if s == "" {
		fmt.Fprint(os.Stderr, "Subscription not set!\n")
		fmt.Fprint(os.Stderr, `Use either the "--subscription"" flag or configure "subscription: ..." in $HOME/.ccv2ctl.yaml\n`)
		os.Exit(1)
	}
	return s
}

func getCookieJar() (j string) {
	j = viper.GetString("cookiejar")
	if j == "" {
		j = findHome()
		j = filepath.Join(j, ".ccv2jar")
	}
	return j
}

func createClient() portal.Client {
	s := getSubscription()
	certPEMBlock, keyPEMBlock := certAndKey()

	return portal.NewClient(s, certPEMBlock, keyPEMBlock, getCookieJar())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ccv2ctl.yaml)")

	rootCmd.PersistentFlags().String("subscription", "", "Subscription to use. Overrides configuration file")
	viper.BindPFlag("subscription", rootCmd.PersistentFlags().Lookup("subscription"))

	rootCmd.PersistentFlags().String("certfile", "", "Path of PEM-encoded client certificate. Overrides value of configuration file")
	viper.BindPFlag("certfile", rootCmd.PersistentFlags().Lookup("certfile"))

	rootCmd.PersistentFlags().String("keyfile", "", "Path of PEM-encoded client certificate key. Overrides value of configuration file")
	viper.BindPFlag("keyfile", rootCmd.PersistentFlags().Lookup("keyfile"))

	rootCmd.PersistentFlags().String("cookiejar", "", "Path of the cookie jar file. Overrides value of configuration file")
	viper.BindPFlag("cookiejar", rootCmd.PersistentFlags().Lookup("cookiejar"))

}

func findHome() string {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return home
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home := findHome()
		viper.AddConfigPath(home)
		viper.SetConfigName(".ccv2ctl")
	}
	viper.SetEnvPrefix("CCV2")
	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
	}
}
