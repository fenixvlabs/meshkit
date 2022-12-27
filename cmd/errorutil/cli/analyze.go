package cli

import "github.com/spf13/cobra"

var analyzeCmd = &cobra.Command{
	Use: "analyze",
	// Aliases: []string{""},
	Short: "analyze",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
