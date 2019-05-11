package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fatih/color"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
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

	w, err := r.Branches()
	failIfErr(err)

	err = w.ForEach(func(arg *plumbing.Reference) error {
		fmt.Println(arg.Name())
		return nil
	})
	failIfErr(err)
}
