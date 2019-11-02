package prompt

import (
	"github.com/c-bata/go-prompt"
)

type completer struct {
	Suggestions []prompt.Suggest
}

func fromRefsToSuggestions(args []ref) []prompt.Suggest {
	ret := make([]prompt.Suggest, len(args))
	for _, ref := range args {
		ret = append(ret, prompt.Suggest{Text: ref.Name().Short()})
	}
	return ret
}

func newCompleter(data []prompt.Suggest) *completer {
	return &completer{
		Suggestions: data,
	}
}
func (c *completer) complete(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(c.Suggestions, d.GetWordBeforeCursor(), true)
}
