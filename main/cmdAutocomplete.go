package main

import prompt "github.com/c-bata/go-prompt"

type cmdAutocomplete interface {
	queueFromCmdLine() string
	msgFromCmdLine() string
}

type cmdAutocompleteImpl struct {
	queuePrompt string
	msgPrompt   string
}

func (ac *cmdAutocompleteImpl) queueFromCmdLine() string {
	queue := prompt.Input(ac.queuePrompt, qCompleter)
	return queue
}

func (ac *cmdAutocompleteImpl) msgFromCmdLine() string {
	message := prompt.Input(ac.msgPrompt, msgCompleter)
	return message
}

func qCompleter(d prompt.Document) []prompt.Suggest {
	suggest := make([]prompt.Suggest, len(sendQueues))
	for i, q := range sendQueues {
		suggest[i] = prompt.Suggest{Text: q, Description: "Send to queue " + q}
	}
	return prompt.FilterHasPrefix(suggest, d.GetWordBeforeCursor(), true)
}

func msgCompleter(d prompt.Document) []prompt.Suggest {
	suggest := []prompt.Suggest{
		{Text: "example", Description: "An example text/plain message"},
	}
	return prompt.FilterHasPrefix(suggest, d.GetWordBeforeCursor(), true)
}
