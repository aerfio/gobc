package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aerfio/gobc/prompt"
	"github.com/fatih/color"
)

var version = "v0.2.3"

func flagInit() (bool, bool, bool) {
	list := flag.Bool("l", false, color.HiRedString("list local and remote branches, without deleting"))
	removeAll := flag.Bool("a", false, color.HiRedString("remove all excess branches, without prompt"))
	printVersion := flag.Bool("v", false, color.HiRedString("show version and exit"))
	flag.Parse()
	return *list, *removeAll, *printVersion
}

func main() {
	list, rmAll, versionFlag := flagInit()

	if versionFlag {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		fmt.Printf("gobc %s from %s", version, exPath)
		os.Exit(0)
	}

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
