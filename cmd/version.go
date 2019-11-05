package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print version information, along with commit HASH and the build date`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version %s, commit %s, date %s", version, commit, date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
