package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
	"fmt"
	message_thrift_service "thrift/source/message/service"
	message_thrift_service_request "thrift/source/message/struct/request"
)

func main() {
	transport, err := thrift.NewTSocketTimeout("192.168.1.57:60000", 3 * time.Second)
	if (err != nil) {
		fmt.Println("socket create error : ", err)
	}
	defer transport.Close()
	if err := transport.Open(); err != nil {
        fmt.Println("transport open error : ", err)
	}
	protocol := thrift.NewTBinaryProtocolTransport(transport)
	mp := thrift.NewTMultiplexedProtocol(protocol, "MessageThriftService")
	ser := message_thrift_service.NewTMessageThriftServiceClientProtocol(transport, mp, mp)

	condition := message_thrift_service_request.NewTSendMessageCondition()
	condition.ID = 123456

	fmt.Println("result : ", ser.SendMessage(condition))
}

