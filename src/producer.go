package main

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
)

var sendQueues = []string{"test.queue.1", "test.queue.2"}

const textPlain string = "text/plain"

func produce() {
	queue := prompt.Input("Queue to send to>>>", qCompleter)
	fmt.Println("Sending to", queue)
}

// func sendMsgs() {
// 	con, er := stomp.Dial("tcp", "localhost:61613")
// 	//defer con.Disconnect()
// 	if er != nil {
// 		fmt.Println("ERROR:", er)
// 	} else {
// 		fmt.Println("Connected successfully")
// 	}
//
// 	for i := 0; i < 100; i++ {
// 		er = con.Send(testSendQ, textPlain, []byte("Hello World "+strconv.Itoa(i)))
// 		if er != nil {
// 			fmt.Println("ERROR:", er)
// 		} else {
// 			fmt.Println("Sent Message")
// 		}
// 	}
// 	con.Disconnect()
// }

func qCompleter(d prompt.Document) []prompt.Suggest {
	suggest := make([]prompt.Suggest, len(sendQueues))
	for i, q := range sendQueues {
		suggest[i] = prompt.Suggest{Text: q, Description: "Send to queue " + q}
	}
	return prompt.FilterHasPrefix(suggest, d.GetWordBeforeCursor(), true)
}
