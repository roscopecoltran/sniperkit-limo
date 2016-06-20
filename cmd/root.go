package cmd

import (
	"fmt"
	"os"

	"github.com/hoop33/limo/config"
	"github.com/hoop33/limo/output"
	"github.com/hoop33/limo/service"
	"github.com/spf13/cobra"
)

var configuration *config.Config
var options struct {
	verbose bool
	output  string
	service string
}

// RootCmd is the root command for limo
var RootCmd = &cobra.Command{
	Use:   "limo",
	Short: "A CLI for managing starred repositories",
	Long: `limo allows you to manage your starred repositories on GitHub, GitLab, and Bitbucket.
You can tag, display, and search your starred repositories.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	flags := RootCmd.PersistentFlags()
	flags.StringVarP(&options.output, "output", "o", "color", "output type")
	flags.StringVarP(&options.service, "service", "s", "github", "service")
	flags.BoolVarP(&options.verbose, "verbose", "v", false, "verbose output")
}

func getConfiguration() (*config.Config, error) {
	if configuration == nil {
		var err error
		if configuration, err = config.ReadConfig(); err != nil {
			return nil, err
		}
	}
	return configuration, nil
}

func getOutput() output.Output {
	return output.ForName(options.output)
}

func getService() (service.Service, error) {
	return service.ForName(options.service)
}
