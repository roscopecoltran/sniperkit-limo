package action

import (
	"fmt"                                             // go-core
	"github.com/roscopecoltran/sniperkit-limo/config" // app-config
	//"github.com/roscopecoltran/sniperkit-limo/model" 										// data-models
	"github.com/sirupsen/logrus" // logs-logrus
	"github.com/spf13/cobra"     // cli-cmd
	//"github.com/davecgh/go-spew/spew" 													// debug-print
	//"github.com/k0kubun/pp" 																// debug-print
)

// AnalyzeCmd does a full-text gateway
var AnalyzeCmd = &cobra.Command{
	Use:     "analyze <vcs uri>",
	Aliases: []string{"analyze", "augmented", "a"},
	Short:   "Analyze info on stars",
	Long:    "Perform an extended analyze on your stars",
	Example: fmt.Sprintf("  %s analyze robust", config.ProgramName),
	Run: func(cmd *cobra.Command, args []string) {

		// do stuff

	},
}

func init() {
	log.WithFields(
		logrus.Fields{
			"prefix":      "app-action",
			"src.file":    "action/analyze.go",
			"cmd.name":    "AnalyzeCmd",
			"method.name": "init()",
			"var.options": options,
		}).Info("registering command...")
	RootCmd.AddCommand(AnalyzeCmd)
}
