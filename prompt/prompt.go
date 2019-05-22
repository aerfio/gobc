package prompt

import (
	"strings"

	unique "github.com/aerfio/gobc/helpers"
	"github.com/aerfio/gobc/types"
	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
)

type ref = types.Ref

func DeletePrompt(toDelete []ref) []string {
	suggestions := fromRefsToSuggestions(toDelete)
	c := NewCompleter(suggestions)
	color.HiGreen("List branches to delete:")
	t := prompt.Input("> ", c.Complete)

	branches := strings.Split(t, " ")

	return unique.UniqStr(branches)
}
