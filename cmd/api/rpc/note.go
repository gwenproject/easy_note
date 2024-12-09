package rpc

import (
	"fmt"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/gwen0x4c3/easy_note/kitex_gen/knote/noteservice"
	"github.com/gwen0x4c3/easy_note/pkg/constants"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var noteClient noteservice.Client

func InitNoteRPC() {
	var err error
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	noteClient, err = noteservice.NewClient(
		constants.NoteServiceName,
		client.WithResolver(r),
		// client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(100*time.Millisecond),   // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
	)
	fmt.Println(noteClient)
	if err != nil {
		panic(err)
	}
}
