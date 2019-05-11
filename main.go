package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func flagInit() bool {
	list := flag.Bool("l", false, color.HiRedString(("list local and remote branches, without deleting")))
	flag.Parse()
	return *list
}

func main() {
	list := flagInit()
	localBranches, remoteBranches := getBranches()
	if list {
		color.Blue("Branches on origin:")
		for _, branch := range remoteBranches {
			fmt.Println(branch.Name().Short())
		}
		color.Magenta("Local branches")
		for _, branch := range localBranches {
			fmt.Println(branch.Name().Short())
		}

		os.Exit(0)
	}

	toDelete := branchesToDelete(localBranches, remoteBranches)

	deleteBranches(toDelete)
	color.HiGreen("Deleted branches:")
	for _, branch := range toDelete {
		color.Green(" - " + branch.Name().Short())
	}
}
