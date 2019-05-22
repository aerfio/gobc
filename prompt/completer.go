package prompt

import (
	"github.com/c-bata/go-prompt"
)

type Completer struct {
	Suggestions []prompt.Suggest
}

func fromRefsToSuggestions(args []ref) []prompt.Suggest {
	var ret []prompt.Suggest
	for _, ref := range args {
		ret = append(ret, prompt.Suggest{Text: ref.Name().Short()})
	}
	return ret
}

func NewCompleter(data []prompt.Suggest) *Completer {
	return &Completer{
		Suggestions: data,
	}
}
func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(c.Suggestions, d.GetWordBeforeCursor(), true)
}
