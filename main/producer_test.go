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

type mockCmdAutocomplete struct {
	queuePrompt    string
	msgPrompt      string
	qPromptCalls   int
	msgPromptCalls int
}

func (ac *mockCmdAutocomplete) queueFromCmdLine() string {
	ac.qPromptCalls++
	return "q"
}

func (ac *mockCmdAutocomplete) msgFromCmdLine() string {
	ac.msgPromptCalls++
	return "msg"
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

func TestProduceGetsParamsFromCommandLine(t *testing.T) {
	amqConn = &amqConnectionMock{protocol: "tcp", address: "localhost:61613", calledCount: 0, calls: make([]sendMsgCallParams, 0)}
	mockAutocomplete := mockCmdAutocomplete{msgPrompt: "What is you message>>> ", queuePrompt: "Queue >>>", msgPromptCalls: 0, qPromptCalls: 0}
	produce(&mockAutocomplete)

	if mockAutocomplete.qPromptCalls != 1 || mockAutocomplete.msgPromptCalls != 1 {
		t.Fail()
	}
}
