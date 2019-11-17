package version

import (
	"fmt"
)

var (
	version             = "dev"
	commit              = "none"
	date                = "unknown"
	PrettyVersionString = fmt.Sprintf("Version %s, commit %s, date %s", version, commit, date)
)
