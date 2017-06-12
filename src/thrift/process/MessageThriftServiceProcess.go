package process

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	message_thrift_service "thrift/source/message/service"
	"service/message"
)

type MessageThriftServiceProcessFactory struct {

}

func NewMessageThriftServiceProcessFactory() *MessageThriftServiceProcessFactory {
	return &MessageThriftServiceProcessFactory{}
}

func (m *MessageThriftServiceProcessFactory) GetServiceProcess() thrift.TProcessor {
	handler := message.NewMessageHandler()
	processor := message_thrift_service.NewTMessageThriftServiceProcessor(handler)
    return processor
}