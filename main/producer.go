package main

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
	"github.com/go-stomp/stomp"
)

var sendQueues = []string{"test.queue.1", "test.queue.2"}

const textPlain string = "text/plain"

type amqConnection interface {
	sendMessage(string, string)
}

type amqConnectionImpl struct {
	protocol string
	address  string
}

func (amqConn *amqConnectionImpl) sendMessage(queue string, message string) {
	con, er := stomp.Dial(amqConn.protocol, amqConn.address)
	defer con.Disconnect()
	if er != nil {
		fmt.Println("ERROR:", er)
		return
	}

	er = con.Send(queue, textPlain, []byte(message))
	if er != nil {
		fmt.Println("ERROR:", er)
		return
	}
}

func produce() {
	queue := prompt.Input("Queue to send to>>> ", qCompleter)
	msg := getMsgFromCmd(queue)
	amqConn := amqConnectionImpl{protocol: "tcp", address: "localhost:61613"}
	sendMsg(queue, msg, &amqConn)
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

func sendMsg(queue string, msg string, amqConn amqConnection) {
	amqConn.sendMessage(queue, msg)
	fmt.Println("Message Sent")
}
