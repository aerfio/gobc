package cmd

import (
	"fmt"

	"github.com/aerfio/gobc/internal/githandler"
	"github.com/spf13/cobra"
	"helm.sh/helm/v3/cmd/helm/require"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "A brief description of your command",

	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Aliases: []string{"delete", "remove"},
	Args:    require.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		local, remote, err := githandler.GetBranches()
		if err != nil {
			return err
		}

		toDelete := githandler.BranchesToDelete(*local, *remote)

		for _, br := range toDelete {
			fmt.Printf("Removing %s...\n", br.Name().Short())
		}

		err = githandler.DeleteBranches(toDelete)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	// git branch -vv | cut -c 3- | awk '$3 !~/\[/ { print $1 }'
	rootCmd.AddCommand(rmCmd)
}
