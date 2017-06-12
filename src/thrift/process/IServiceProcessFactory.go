package process

import "git.apache.org/thrift.git/lib/go/thrift"

type IServiceProcessFactory interface {
	GetServiceProcess() thrift.TProcessor
}

