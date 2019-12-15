package list

import (
	"fmt"

	"github.com/aerfio/gobc/internal/githandler"
	"github.com/jedib0t/go-pretty/list"
)

func getBranches() (local, remote []interface{}, err error) {
	localRef, remoteRef, err := githandler.GetBranches()

	if err != nil {
		return nil, nil, err
	}

	localBranches, remoteBranches := make([]interface{}, 0), make([]interface{}, 0)
	for _, loc := range *localRef {
		localBranches = append(localBranches, loc.Name().Short())
	}
	for _, rem := range *remoteRef {
		remoteBranches = append(remoteBranches, rem.Name().Short())
	}
	return localBranches, remoteBranches, nil
}

func Print() error {
	local, remote, err := getBranches()
	if err != nil {
		return err
	}
	l := list.NewWriter()
	l.SetStyle(list.StyleDefault)
	l.Indent()
	l.AppendItem("local")
	l.Indent()
	l.AppendItems(local)
	l.UnIndent()
	l.AppendItem("origin")
	l.Indent()
	l.AppendItems(remote)
	fmt.Println("Branches:")
	fmt.Println(l.Render())
	return nil
}
