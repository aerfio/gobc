package main

import (
	"flag"
	"os"

	"github.com/fatih/color"
)

func flagInit() bool {
	list := flag.Bool("l", false, color.HiRedString("list local and remote branches, without deleting"))
	flag.Parse()
	return *list
}

func main() {
	localBranches, remoteBranches := getBranches()
	if flagInit() {
		listBranches(localBranches, remoteBranches)
		os.Exit(0)
	}

	toDelete := branchesToDelete(localBranches, remoteBranches)

	deleteBranches(toDelete)
	color.HiGreen("Deleted branches:")
	for _, branch := range toDelete {
		color.Green(" - " + branch.Name().Short())
	}
}
