package main

import (
	"fmt"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/gwen0x4c3/easy_note/cmd/user/dal"
	"github.com/gwen0x4c3/easy_note/kitex_gen/kuser/userservice"
	"github.com/gwen0x4c3/easy_note/pkg/constants"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	// TODO Add Jaeger

	// TODO Init NoteService rpc client

	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	ip, err := constants.GetOutBoundIP()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ip, constants.UserServicePort))
	if err != nil {
		panic(err)
	}
	Init()
	svr := userservice.NewServer(
		new(UserServiceImpl),
		server.WithServiceAddr(addr), // address
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}),
		server.WithRegistry(r), // registry
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(), // multiplex
	)

	err = svr.Run()

	if err != nil {
		klog.Info(err.Error())
	}
}
