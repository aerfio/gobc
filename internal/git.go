package internal

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"go.uber.org/zap"
	"path/filepath"
)

func parentPath(depth int) string {
	if depth <= 0 {
		return "."
	}
	parents := make([]string, 0)
	for i := 0; i < depth; i++ {
		parents = append(parents, "..")
	}

	return filepath.Join(parents...)
}

func openRepo(depth int) (*git.Repository, error) {
	for i := 0; i < depth; i++ {
		r, err := git.PlainOpen(parentPath(i))
		if err != nil {
			if errors.Is(err, git.ErrRepositoryNotExists) {
				continue
			}
			return nil, err
		}
		return r, nil
	}
	return nil, git.ErrRepositoryNotExists
}

func getRemoteBranches(repo *git.Repository, remoteName string) ([]*plumbing.Reference, error) {
	remote, err := repo.Remote(remoteName)
	if err != nil {
		return nil, err
	}

	refs, err := remote.List(&git.ListOptions{})
	if err != nil {
		return nil, err
	}

	branches := make([]*plumbing.Reference, 0)
	for _, ref := range refs {
		if ref.Name().IsBranch() {
			branches = append(branches, ref)
		}
	}
	return branches, nil
}

func getLocalBranches(repo *git.Repository) ([]*plumbing.Reference, error) {
	brIterator, err := repo.Branches()
	if err != nil {
		return nil, err
	}
	defer brIterator.Close()

	branches := make([]*plumbing.Reference, 0)
	err = brIterator.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().IsBranch() {
			branches = append(branches, ref)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return branches, nil
}

func branchesToDelete(localBranches, remoteBranches []*plumbing.Reference) []*plumbing.Reference {
	toDelete := make([]*plumbing.Reference, 0)

	for _, local := range localBranches {
		if local.Name().Short() == "master" || local.Name().Short() == "main" {
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

func deleteBranches(logger *zap.SugaredLogger, r *git.Repository, branches []*plumbing.Reference) error {
	for _, branch := range branches {
		logger.Debugw("removing branch", "name", branch.Name().Short())
		err := r.Storer.RemoveReference(branch.Name())
		if err != nil {
			return fmt.Errorf("while deleting branch %s: %w", branch.Name().String(), err)
		}
	}

	return nil
}
