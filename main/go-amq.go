package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

const prodAction string = "Produce"
const readAction string = "Consume"

func main() {
	action := prompt.Input("Action>>>", rwCompleter)

	switch action {
	case prodAction:
		produce()
	case readAction:
		consume()
	default:
		fmt.Println("ERROR: Invalid Action")
	}
}

func rwCompleter(d prompt.Document) []prompt.Suggest {
	suggest := []prompt.Suggest{
		{Text: prodAction, Description: "Send some messages"},
		{Text: readAction, Description: "Read some messages"},
	}
	return prompt.FilterHasPrefix(suggest, d.GetWordBeforeCursor(), true)
}
