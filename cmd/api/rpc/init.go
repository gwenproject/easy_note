package rpc

import (
	"github.com/gwen0x4c3/easy_note/pkg/errno"
)

func InitRPC() {
	InitUserRPC()
	InitNoteRPC()
}

func checkError(code int64, message string, err error) error {
	if err != nil {
		return err
	}
	if code != errno.SuccessCode {
		return errno.NewErrNo(code, message)
	}
	return nil
}
