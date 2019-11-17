package cmd

import (
	"github.com/aerfio/gobc/internal/version"
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/cmd/helm/require"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print version information, along with commit HASH and the build date`,
	Args:  require.NoArgs,
	Run:   version.PrettyVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
