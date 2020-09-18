package main

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
	"github.com/go-stomp/stomp"
)

var sendQueues = []string{"test.queue.1", "test.queue.2"}

const textPlain string = "text/plain"

var stompDial = stomp.Dial

func produce() {
	queue := prompt.Input("Queue to send to>>> ", qCompleter)
	msg := getMsgFromCmd(queue)
	sendMsg(queue, msg)
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

func getMsgFromCmd(queue string) string {
	fmt.Println("Sending to", queue)
	msg := prompt.Input("What is your message?>>> ", msgCompleter)

	return msg
}

func sendMsg(queue string, msg string) {
	con, er := stompDial("tcp", "localhost:61613")
	defer con.Disconnect()
	if er != nil {
		fmt.Println("ERROR:", er)
		return
	}

	er = con.Send(queue, textPlain, []byte(msg))
	if er != nil {
		fmt.Println("ERROR:", er)
	}
	fmt.Println("Message Sent")
}
