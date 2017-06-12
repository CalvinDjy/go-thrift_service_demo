package boot

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"time"
	"config"
	"flag"
)

type Server struct {

}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {
	confFile := flag.String("C", "src/config/thrift.ini", "config file path")
	flag.Parse();
	thriftConfig, err := NewConfig(confFile)
	if err != nil {
		fmt.Println("Config error : ", err)
		return
	}
	if !thriftConfig.Has("thrift::ip") || !thriftConfig.Has("thrift::port") {
		fmt.Println("Config error : can not find server ip or port")
		return
	}

	serverAddr := thriftConfig.String("thrift::ip") + ":" + thriftConfig.String("thrift::port")
	serverTransport, err := thrift.NewTServerSocketTimeout(serverAddr, 3 * time.Second)
	if err != nil {
		fmt.Println("Init socket error : ", err)
		return
	}

	multiProcess := thrift.NewTMultiplexedProcessor()
	for name, processFactory := range config.Process_factory_list {
		multiProcess.RegisterProcessor(name, processFactory.GetServiceProcess())
		fmt.Println("[***** Register processor [", name, "] *****]")
	}

	server := thrift.NewTSimpleServer2(multiProcess, serverTransport)
	fmt.Println("Starting server... on ", serverAddr)
	fmt.Println(server.Serve())
}