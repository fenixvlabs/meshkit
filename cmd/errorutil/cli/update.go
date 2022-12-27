package cli

import "github.com/spf13/cobra"

var updateCmd = &cobra.Command{
	Use: "update",
	// Aliases: []string{""},
	Short: "update",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
