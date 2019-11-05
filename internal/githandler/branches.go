package githandler

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type ref = *plumbing.Reference

func failIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func deleteBranches(toDelete []ref) {
	r, err := git.PlainOpen(".")
	failIfErr(err)
	for _, branch := range toDelete {
		err := r.Storer.RemoveReference(branch.Name())
		failIfErr(err)
	}
}

func BranchToString(b ref) string {
	return b.Name().Short()
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

func GetBranches() ([]ref, []ref) {
	r, err := git.PlainOpen(".")
	failIfErr(err)

	remote, err := r.Remote("origin")
	failIfErr(err)

	remoteBranches := make([]ref, 0)

	refs, err := remote.List(&git.ListOptions{})
	failIfErr(err)
	for _, ref := range refs {
		if ref.Name().IsBranch() {
			remoteBranches = append(remoteBranches, ref)
		}
	}

	w, err := r.Branches()
	failIfErr(err)

	localBranches := make([]ref, 0)
	err = w.ForEach(func(arg ref) error {
		if arg.Name().IsBranch() {
			localBranches = append(localBranches, arg)
		}
		return nil
	})
	failIfErr(err)

	return localBranches, remoteBranches
}

func delBranchesFromStr(branches []string) {
	localBranches, _ := GetBranches()

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
