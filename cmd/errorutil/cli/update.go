package cli

import (
	"github.com/fenixvlabs/meshkit/pkg/errorutil"
	"github.com/spf13/cobra"
)

var updateAll bool

func updateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update error codes and details",
		Long:  "Update releases error codes where specified, and updates error details",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			globalFlags, err := getGlobalFlags(cmd)
			if err != nil {
				return err
			}
			updateAll, err := cmd.Flags().GetBool(forceUpdateAllCodesCmdFlag)
			if err != nil {
				return err
			}
			return errorutil.ExportWalk(globalFlags, true, updateAll)
		},
	}
	cmd.PersistentFlags().BoolVar(&updateAll, forceUpdateAllCodesCmdFlag, false, "Update and re-sequence all error codes.")
	return cmd
}

func init() {
	// rootCmd.AddCommand(updateCmd)
}
