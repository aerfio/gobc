package githandler

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

type ref = *plumbing.Reference

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

func BranchesToDelete(localBranches []ref, remoteBranches []ref) []ref {
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

func GetBranches() (*[]ref, *[]ref, error) {
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

// func delBranchesFromStr(branches []string) {
// 	localBranches, _ := GetBranches()
//
// 	toDel := make([]ref, 0)
//
// 	for _, refBranch := range localBranches {
// 		for _, branch := range branches {
// 			if refBranch.Name().Short() == branch {
// 				toDel = append(toDel, refBranch)
// 			}
// 		}
// 	}
// 	deleteBranches(toDel)
// }
