package version

import (
	"fmt"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	// PrettyVersionString is used during release process to be overridden using ldflags
	PrettyVersionString = fmt.Sprintf("Version %s, commit %s, date %s", version, commit, date)
)
