package config

import "thrift/process"

var Process_factory_list = map[string]process.IServiceProcessFactory {
	"MessageThriftService" : process.NewMessageThriftServiceProcessFactory(),
}