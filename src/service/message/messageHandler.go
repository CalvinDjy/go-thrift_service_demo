package message

import (
	"fmt"
	"thrift/source/message/struct/request"
)

type MessageHandler struct {
	a int
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		a: 123,
	}
}

func (p *MessageHandler) SendMessage(sendMessageCondition *request.TSendMessageCondition) (err error) {
	fmt.Println(p.a)
	fmt.Println(sendMessageCondition.GetID())
	return nil
}