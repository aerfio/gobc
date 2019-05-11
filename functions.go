package main

import (
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

func getBranches() ([]ref, []ref) {
	r, err := git.PlainOpen(".")
	failIfErr(err)

	remotes, err := r.Remotes()
	failIfErr(err)
	remoteBranches := make([]ref, 0)

	for _, remote := range remotes {
		refs, err := remote.List(&git.ListOptions{})
		failIfErr(err)
		for _, ref := range refs {
			if ref.Name().IsBranch() {
				remoteBranches = append(remoteBranches, ref)
			}
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
