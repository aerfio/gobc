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
	localBranches, remoteBranches := getBranches()
	list, rmAll := flagInit()
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
		deleteBranches(toDelete)
		color.HiGreen("Deleted branches:")
		for _, branch := range toDelete {
			color.Green(" - " + branch.Name().Short())
		}
	}
	chosenBranches := prompt.DeletePrompt(toDelete)
	delBranchesFromStr(chosenBranches)
	color.Green("Success!")
}
