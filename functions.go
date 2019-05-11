package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func failIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type ref = *plumbing.Reference

func deleteBranches(toDelete []ref) {
	if len(toDelete) == 0 {
		color.Green("There's nothing to delete!")
		os.Exit(0)
	}

	r, err := git.PlainOpen(".")
	failIfErr(err)
	for _, branch := range toDelete {
		err := r.Storer.RemoveReference(branch.Name())
		failIfErr(err)
	}
}

func listBranches(localBranches []ref, remoteBranches []ref){
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

func getBranches() ([]ref, []ref) {
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
