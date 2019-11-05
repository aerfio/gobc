package cmd

import (
	`os`

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates zsh completion scripts",
	Long: `To load completion run

. <(gobc completion)

To configure your bash shell to load completions for each session add to your .zshrc

. <(gobc completion)
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return rootCmd.GenZshCompletion(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
