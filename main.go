package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fatih/color"
	"gopkg.in/src-d/go-git.v4"
)

func flagInit() (int, string) {
	color.Set(color.Bold)
	prID := flag.Int("n", -1, color.HiRedString("number of pull request you want to fetch for review - mandatory"))
	branch := flag.String("b", "review", color.HiBlueString("name of branch you want PR fetched to"))
	color.Unset()
	flag.Parse()
	return *prID, *branch
}

func failIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// number, branch := flagInit()
	// if number < 0 {
	// 	log.Fatal(fmt.Errorf(color.RedString(`use "-n" flag to set number of pull request you want to fetch`)))

	// }
	r, err := git.PlainOpen(".")
	failIfErr(err)

	remotes, err := r.Remotes()
	failIfErr(err)

	for _, remote := range remotes {
		// fmt.Println(remote.List(&git.ListOptions{}))
		refs, err := remote.List(&git.ListOptions{})
		failIfErr(err)
		for _, ref := range refs {
			fmt.Println(ref)
		}
	}

	// w, err := r.Branches()
	// failIfErr(err)

	// err = w.ForEach(func(arg *plumbing.Reference) error {
	// 	fmt.Println(arg.Name())
	// 	return nil
	// })
	// failIfErr(err)
	// for i, num :=range w {
	// 	fmt.Println()
	// }
	// externalRefs := config.RefSpec(fmt.Sprintf("refs/pull/%d/head:refs/heads/%s", number, branch))
	// err = r.Fetch(&git.FetchOptions{Progress: nil, RemoteName: "upstream", RefSpecs: []config.RefSpec{externalRefs}})
	// checkError(err)

	// w, err := r.Worktree()
	// checkError(err)

	// branchAsPlmbRef := plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch))
	// err = w.Checkout(&git.CheckoutOptions{Branch: branchAsPlmbRef})

	// checkError(err)
	// color.HiGreen("Done!")
}
