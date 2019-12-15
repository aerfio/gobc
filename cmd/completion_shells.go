package cmd

import (
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/cmd/helm/require"
)

//
// import (
// 	"fmt"
//
// 	"github.com/spf13/cobra"
// )
//
// // zshCmd represents the zsh command
// var zshCmd = &cobra.Command{
// 	Use:   "zsh",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:
//
// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("zsh called")
// 	},
// }
//
// func init() {
// 	completionCmd.AddCommand(zshCmd)
//
// 	// Here you will define your flags and configuration settings.
//
// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// zshCmd.PersistentFlags().String("foo", "", "A help for foo")
//
// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// zshCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

func init() {
	completionCmd.AddCommand(zshCmd)
	completionCmd.AddCommand(bashCmd)
}

var zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: `Output zsh completion script for gobc.`,
	Long: `
Generates a zsh autocompletion script for gobc.

To load completion run

	. <(gobc completion zsh)
`,
	Args: require.NoArgs,
	RunE: func(command *cobra.Command, args []string) error {
		return rootCmd.GenZshCompletion(command.OutOrStdout())
	},
}

var bashCmd = &cobra.Command{
	Use:   "bash",
	Short: `Output bash completion script for gobc.`,
	Long: `
Generates a bash autocompletion script for gobc.

To load completion run

	. <(gobc completion bash)
`,
	Args: require.NoArgs,
	RunE: func(command *cobra.Command, args []string) error {
		return rootCmd.GenBashCompletion(command.OutOrStdout())
	},
}
