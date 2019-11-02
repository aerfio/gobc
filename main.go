package main

import (
	"flag"
	"os"

	"github.com/aerfio/gobc/prompt"
	"github.com/fatih/color"
)

func flagInit() (bool, bool) {
	list := flag.Bool("l", false, color.HiRedString("list local and remote branches, without deleting"))
	removeAll := flag.Bool("a", false, color.HiRedString("remove all excess branches, without prompt"))
	flag.Parse()
	return *list, *removeAll
}

func main() {
	list, rmAll := flagInit()
	localBranches, remoteBranches := getBranches()
	toDelete := branchesToDelete(localBranches, remoteBranches)

	if list {
		listBranches(localBranches, remoteBranches)
		printExcess(toDelete)
		os.Exit(0)
	}
	if len(toDelete) == 0 {
		color.Green("There's nothing to delete!")
		os.Exit(0)
	}
	if rmAll {
		color.HiGreen("Deleted branches:")
		for _, branch := range toDelete {
			color.Green(" - " + branch.Name().Short())
		}
		deleteBranches(toDelete)
		os.Exit(0)
	}
	chosenBranches := prompt.DeletePrompt(toDelete)
	delBranchesFromStr(chosenBranches)
	color.Green("Success!")
}
