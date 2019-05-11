package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func flagInit() bool {
	list := flag.Bool("l", false, color.HiRedString(("only list local and remote branches")))
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

	toDelete := make([]ref, 0)

	for _, local := range localBranches {
		if local.Name().Short() == "master" {
			continue
		}

		rm := true
		for _, remote := range remoteBranches {
			if local.Name().Short() == remote.Name().Short() {
				rm = false
			}
		}

		if rm {
			toDelete = append(toDelete, local)
		}
	}

	deleteBranches(toDelete)
	color.HiGreen("Deleted branches:")
	for _, branch := range toDelete {
		color.Green(" - " + branch.Name().Short())
	}
}
