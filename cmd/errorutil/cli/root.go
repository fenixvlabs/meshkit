package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version = "0.0.1"
var verbose = "verbose"
var rootDir = "dir"
var outDir = "out-dir"
var infoDir = "info-dir"
var skipDirs = "skip-dirs"
var forceUpdateAllCodesCmdFlag = "force"

var rootCmd = &cobra.Command{
	Use:     "errorutil",
	Version: version,
	Short:   "errorutil",
	Long:    "errorutil",

	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// rootCmd := &cobra.Command{Use: errorutil.App}
	rootCmd.PersistentFlags().BoolP(verbose, "v", false, "verbose output")
	rootCmd.PersistentFlags().StringP(rootDir, "d", ".", "root directory")
	rootCmd.PersistentFlags().StringP(outDir, "o", "", "output directory")
	rootCmd.PersistentFlags().StringP(infoDir, "i", "", "directory containing the component_info.json file")
	rootCmd.PersistentFlags().StringSlice(skipDirs, []string{}, "directories to skip (comma-separated list, repeatable argument)")
	rootCmd.AddCommand(updateCmd())
}

func initConfig() {
	return
}
