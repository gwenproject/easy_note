package main

import (
    "fmt"
    "github.com/gwen0x4c3/easy_note/cmd/note/rpc"
    "github.com/gwen0x4c3/easy_note/pkg/tracer"
    "net"

    "github.com/cloudwego/kitex/pkg/klog"
    "github.com/cloudwego/kitex/pkg/limit"
    "github.com/cloudwego/kitex/pkg/rpcinfo"
    "github.com/cloudwego/kitex/server"
    "github.com/gwen0x4c3/easy_note/cmd/note/dal"
    "github.com/gwen0x4c3/easy_note/kitex_gen/knote/noteservice"
    "github.com/gwen0x4c3/easy_note/pkg/constants"
    etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
    tracer.InitJaeger(constants.NoteServiceName)

    rpc.InitRPC()

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
    addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ip, constants.NoteServicePort))
    if err != nil {
        panic(err)
    }
    Init()
    svr := noteservice.NewServer(
        new(NoteServiceImpl),
        server.WithServiceAddr(addr),
        server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.NoteServiceName}),
        server.WithRegistry(r),
        server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
        //server.WithMuxTransport(),
    )

    err = svr.Run()

    if err != nil {
        klog.Info(err.Error())
    }
}
