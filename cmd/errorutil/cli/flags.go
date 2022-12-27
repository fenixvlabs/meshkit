package cli

import (
	"github.com/fenixvlabs/meshkit/internal/errorutil"
	"github.com/spf13/cobra"
)

func defaultIfEmpty(value, defaultValue string) string {
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func getGlobalFlags(cmd *cobra.Command) (errorutil.GlobalFlags, error) {
	flags := errorutil.GlobalFlags{}
	verbose, err := cmd.Flags().GetBool(verbose)
	if err != nil {
		return flags, err
	}
	flags.Verbose = verbose
	rootDir, err := cmd.Flags().GetString(rootDir)
	if err != nil {
		return flags, err
	}
	flags.RootDir = rootDir
	skipDirs, err := cmd.Flags().GetStringSlice(skipDirs)
	if err != nil {
		return flags, err
	}
	flags.SkipDirs = skipDirs
	outDir, err := cmd.Flags().GetString(outDir)
	if err != nil {
		return flags, err
	}
	flags.OutDir = defaultIfEmpty(outDir, rootDir)
	infoDir, err := cmd.Flags().GetString(infoDir)
	if err != nil {
		return flags, err
	}
	flags.InfoDir = defaultIfEmpty(infoDir, rootDir)
	return flags, nil
}
