package cmd

import (
	"log"
	"os"

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
}

var zshCmd = &cobra.Command{
	Use:   "zsh [output_file]",
	Short: `Output zsh completion script for gobc.`,
	Long: `
Generates a zsh autocompletion script for gobc.

This writes to /usr/share/zsh/vendor-completions/_gobc by default so will
probably need to be run with sudo or as root, eg

    sudo gobc genautocomplete zsh

Logout and login again to use the autocompletion scripts, or source
them directly

    autoload -U compinit && compinit

If you supply a command line argument the script will be written
there.
`,
	Args: require.NoArgs,
	Run: func(command *cobra.Command, args []string) {
		out := "/usr/share/zsh/vendor-completions/_gobc"
		outFile, err := os.Create(out)
		if err != nil {
			log.Fatal(err)
		}
		defer func() { _ = outFile.Close() }()
		err = rootCmd.GenZshCompletion(outFile)
		if err != nil {
			log.Fatal(err)
		}
	},
}
