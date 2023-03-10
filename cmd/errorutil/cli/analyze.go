package cli

import (
	"github.com/fenixvlabs/meshkit/internal/errorutil"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze a directory tree",
	Long:  "analyze analyzes a directory tree for error codes ",
	Args:  cobra.MinimumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		globalFlags, err := getGlobalFlags(cmd)
		if err != nil {
			return err
		}
		return errorutil.ExportWalk(globalFlags, false, false)
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
