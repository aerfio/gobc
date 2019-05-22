package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
)

type Suggestion struct {
	Text string
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {

}
func deletePrompt(toDelete []ref) {
	color.Green("List branches to delete:")
	t := prompt.Input("> ", completer)
	fmt.Println("You selected " + t)
}
