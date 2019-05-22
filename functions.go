package main

import (
	"fmt"
	"log"

	"github.com/aerfio/gobc/types"
	"github.com/fatih/color"
	"gopkg.in/src-d/go-git.v4"
)

type Ref = types.Ref

func failIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func deleteBranches(toDelete []Ref) {
	r, err := git.PlainOpen(".")
	failIfErr(err)
	for _, branch := range toDelete {
		err := r.Storer.RemoveReference(branch.Name())
		failIfErr(err)
	}
}

func listBranches(localBranches []Ref, remoteBranches []Ref) {
	color.Blue("Branches on origin:")
	for _, branch := range remoteBranches {
		fmt.Println(branch.Name().Short())
	}
	color.Magenta("Local branches")
	for _, branch := range localBranches {
		fmt.Println(branch.Name().Short())
	}
}

func branchesToDelete(localBranches []Ref, remoteBranches []Ref) []Ref {
	toDelete := make([]Ref, 0)

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
	return toDelete
}

func printExcess(refs []Ref) {
	color.Yellow("Excess branches:")
	for _, branch := range refs {
		fmt.Println(branch.Name().Short())
	}
}

func getBranches() ([]Ref, []Ref) {
	r, err := git.PlainOpen(".")
	failIfErr(err)

	remote, err := r.Remote("origin")
	failIfErr(err)

	remoteBranches := make([]Ref, 0)

	refs, err := remote.List(&git.ListOptions{})
	failIfErr(err)
	for _, ref := range refs {
		if ref.Name().IsBranch() {
			remoteBranches = append(remoteBranches, ref)
		}
	}

	w, err := r.Branches()
	failIfErr(err)

	localBranches := make([]Ref, 0)
	err = w.ForEach(func(arg Ref) error {
		if arg.Name().IsBranch() {
			localBranches = append(localBranches, arg)
		}
		return nil
	})
	failIfErr(err)

	return localBranches, remoteBranches
}

func delBranchesFromStr(branches []string) {
	localBranches, _ := getBranches()

	toDel := make([]Ref, 0)

	for _, refBranch := range localBranches {
		for _, branch := range branches {
			if refBranch.Name().Short() == branch {
				toDel = append(toDel, refBranch)
			}
		}
	}
	deleteBranches(toDel)
}
