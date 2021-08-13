package internal

import (
	"fmt"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	// prettyVersionString is used during release process to be overridden using ldflags.
	prettyVersionString = fmt.Sprintf("Version %s, commit %s, date %s", version, commit, date)
)
