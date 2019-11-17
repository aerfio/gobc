package list

import (
	"fmt"

	"github.com/aerfio/gobc/internal/githandler"
	"github.com/jedib0t/go-pretty/list"
)

func getBranches() (local, remote []interface{}) {
	localRef, remoteRef := githandler.GetBranches()
	localBranches, remoteBranches := make([]interface{}, 0), make([]interface{}, 0)
	for _, loc := range localRef {
		localBranches = append(localBranches, githandler.BranchToString(loc))
	}
	for _, rem := range remoteRef {
		remoteBranches = append(remoteBranches, githandler.BranchToString(rem))
	}
	return localBranches, remoteBranches
}

func Print() {
	local, remote := getBranches()
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
}
