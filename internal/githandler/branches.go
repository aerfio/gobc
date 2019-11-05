package githandler

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type ref = *plumbing.Reference

func FailIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func deleteBranches(toDelete []ref) {
	r, err := git.PlainOpen(".")
	FailIfErr(err)
	for _, branch := range toDelete {
		err := r.Storer.RemoveReference(branch.Name())
		FailIfErr(err)
	}
}

func listBranches(localBranches []ref, remoteBranches []ref) {
	color.Blue("Branches on origin:")
	for _, branch := range remoteBranches {
		fmt.Println(branch.Name().Short())
	}
	color.Magenta("Local branches")
	for _, branch := range localBranches {
		fmt.Println(branch.Name().Short())
	}
}

func branchesToDelete(localBranches []ref, remoteBranches []ref) []ref {
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
	return toDelete
}

func printExcess(refs []ref) {
	if len(refs) == 0 {
		color.Yellow("There's no excess branches!")
		return
	}
	color.Yellow("Excess branches:")
	for _, branch := range refs {
		fmt.Println(branch.Name().Short())
	}
}

func getBranches() ([]ref, []ref) {
	r, err := git.PlainOpen(".")
	FailIfErr(err)

	remote, err := r.Remote("origin")
	FailIfErr(err)

	remoteBranches := make([]ref, 0)

	refs, err := remote.List(&git.ListOptions{})
	FailIfErr(err)
	for _, ref := range refs {
		if ref.Name().IsBranch() {
			remoteBranches = append(remoteBranches, ref)
		}
	}

	w, err := r.Branches()
	FailIfErr(err)

	localBranches := make([]ref, 0)
	err = w.ForEach(func(arg ref) error {
		if arg.Name().IsBranch() {
			localBranches = append(localBranches, arg)
		}
		return nil
	})
	FailIfErr(err)

	return localBranches, remoteBranches
}

func delBranchesFromStr(branches []string) {
	localBranches, _ := getBranches()

	toDel := make([]ref, 0)

	for _, refBranch := range localBranches {
		for _, branch := range branches {
			if refBranch.Name().Short() == branch {
				toDel = append(toDel, refBranch)
			}
		}
	}
	deleteBranches(toDel)
}
