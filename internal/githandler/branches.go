package githandler

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type ref = *plumbing.Reference

// DeleteBranches deletes branches provided by argument. For now only from current folder
func DeleteBranches(toDelete []ref) error {
	r, err := git.PlainOpen(".")
	if err != nil {
		return err
	}

	for _, branch := range toDelete {
		err := r.Storer.RemoveReference(branch.Name())
		if err != nil {
			return err
		}
	}

	return nil
}

// BranchesToDelete lists excess branches, that are not on remote anymore but are in your local git repo
func BranchesToDelete(localBranches, remoteBranches []ref) []ref {
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

// GetBranches lists local and remote branches (only origin)
func GetBranches() (localBr, remoteBr *[]ref, err error) {
	r, err := git.PlainOpen(".")
	if err != nil {
		return nil, nil, err
	}

	remote, err := r.Remote("origin")
	if err != nil {
		return nil, nil, err
	}

	remoteBranches := make([]ref, 0)

	refs, err := remote.List(&git.ListOptions{})
	if err != nil {
		return nil, nil, err
	}

	for _, ref := range refs {
		if ref.Name().IsBranch() {
			remoteBranches = append(remoteBranches, ref)
		}
	}

	w, err := r.Branches()
	if err != nil {
		return nil, nil, err
	}

	localBranches := make([]ref, 0)
	err = w.ForEach(func(arg ref) error {
		if arg.Name().IsBranch() {
			localBranches = append(localBranches, arg)
		}
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return &localBranches, &remoteBranches, nil
}
