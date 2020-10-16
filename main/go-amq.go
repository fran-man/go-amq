package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

const prodAction string = "Produce"
const readAction string = "Consume"
const exitAction string = "Exit"

var realAutocomplete cmdAutocompleteImpl

func main() {
	amqConn = &amqConnectionImpl{protocol: "tcp", address: "localhost:61613"}
	for {
		action := prompt.Input("Action>>> ", rwCompleter)

		realAutocomplete = cmdAutocompleteImpl{msgPrompt: "What is you message>>> ", queuePrompt: "Queue >>>"}

		switch action {
		case prodAction:
			produce(&realAutocomplete)
		case readAction:
			consume()
		case exitAction:
			return
		default:
			fmt.Println("ERROR: Invalid Action")
		}
	}
}

func rwCompleter(d prompt.Document) []prompt.Suggest {
	suggest := []prompt.Suggest{
		{Text: prodAction, Description: "Send some messages"},
		{Text: readAction, Description: "Read some messages"},
		{Text: exitAction, Description: "Close the program"},
	}
	return prompt.FilterHasPrefix(suggest, d.GetWordBeforeCursor(), true)
}
