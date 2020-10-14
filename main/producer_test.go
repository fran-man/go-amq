package main

import (
	"testing"
)

type amqConnectionMock struct {
	protocol    string
	address     string
	calledCount int
	calls       []sendMsgCallParams
}

type sendMsgCallParams struct {
	queue   string
	message string
}

func (amqConn *amqConnectionMock) sendMessage(queue string, message string) {
	amqConn.calledCount++
	params := sendMsgCallParams{queue: queue, message: message}
	amqConn.calls = append(amqConn.calls, params)
}

func TestSendMessageIsCalled(t *testing.T) {
	amqConnMock := amqConnectionMock{protocol: "tcp", address: "localhost:61613", calledCount: 0, calls: make([]sendMsgCallParams, 0)}
	sendMsg("q", "msg", &amqConnMock)

	if amqConnMock.calledCount != 1 {
		t.Fail()
	}

	if amqConnMock.calls[0].queue != "q" || amqConnMock.calls[0].message != "msg" {
		t.Fail()
	}
}
