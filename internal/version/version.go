package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func PrettyVersion(cmd *cobra.Command, args []string) {
	fmt.Printf(PrettyString())
}

func PrettyString() string {
	return fmt.Sprintf("Version %s, commit %s, date %s", version, commit, date)
}
