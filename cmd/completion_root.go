package cmd

import (
	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [shell]",
	Short: `Output completion script for a given shell.`,
	Long: `
Generates a shell completion script for gobc.
Run with --help to list the supported shells.
`,
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
