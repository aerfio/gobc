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
	c := newCompleter(suggestions)
	color.HiGreen("List branches to delete:")
	t := prompt.Input("> ", c.Complete, prompt.OptionPrefixTextColor(prompt.Yellow),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray))

	branches := strings.Split(t, " ")

	return unique.UniqStr(branches)
}
