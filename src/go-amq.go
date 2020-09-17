package main

import (
	"fmt"
	"strconv"

	"github.com/go-stomp/stomp"
)

const testSendQ string = "test.queue.1"
const textPlain string = "text/plain"

func main() {
	con, er := stomp.Dial("tcp", "localhost:61613")
	//defer con.Disconnect()
	if er != nil {
		fmt.Println("ERROR:", er)
	} else {
		fmt.Println("Connected successfully")
	}

	for i := 0; i < 100; i++ {
		er = con.Send(testSendQ, textPlain, []byte("Hello World "+strconv.Itoa(i)))
		if er != nil {
			fmt.Println("ERROR:", er)
		} else {
			fmt.Println("Sent Message")
		}
	}
	con.Disconnect()
}
