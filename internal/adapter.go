package internal

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func deleteStaleBranches(sugared *zap.SugaredLogger, ctx *cli.Context) error {
	if ctx.Bool("version") {
		fmt.Println(prettyVersionString)
		return nil
	}

	depth := ctx.Uint("depth")
	verbose := ctx.Bool("verbose")
	remoteName := ctx.String("remote")

	var opts []zap.Option
	if !verbose {
		opts = append(opts, zap.IncreaseLevel(zapcore.InfoLevel))
	}

	log := sugared.Desugar().WithOptions(opts...).Sugar()
	defer log.Sync()

	log.Debugw("flag values", "depth", depth, "verbose", verbose, "remote", remoteName)

	log.Debug("opening repo")
	repo, err := openRepo(int(depth))
	if err != nil {
		return fmt.Errorf("while opening git repo: %w", err)
	}
	log.Debug("listing local branches")
	local, err := getLocalBranches(repo)
	if err != nil {
		return fmt.Errorf("while listing local branches: %w", err)
	}

	log.Debug("listing remote branches")
	remote, err := getRemoteBranches(repo, remoteName)
	if err != nil {
		return fmt.Errorf("while listing remote branches: %w", err)
	}

	log.Debug("creating a list of branches to delete")
	toDelete := branchesToDelete(local, remote)

	log.Debug("deleting stale branches")
	return deleteBranches(log, repo, toDelete)
}

func DeleteStaleBranchesWithExitCode(sugared *zap.SugaredLogger) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		err := deleteStaleBranches(sugared, ctx)
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}
		return nil
	}
}
