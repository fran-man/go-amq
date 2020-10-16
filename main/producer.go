package main

import (
	"fmt"

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

var amqConn amqConnection

func produce(autocomplete cmdAutocomplete) {
	queue := autocomplete.queueFromCmdLine()
	msg := autocomplete.msgFromCmdLine()
	sendMsg(queue, msg, amqConn)
}

func sendMsg(queue string, msg string, amqConn amqConnection) {
	amqConn.sendMessage(queue, msg)
	fmt.Println("Message Sent")
}
