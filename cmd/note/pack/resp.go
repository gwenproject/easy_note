package pack

import (
    "errors"
    "github.com/gwen0x4c3/easy_note/kitex_gen/knote"
    "github.com/gwen0x4c3/easy_note/pkg/errno"
)

func BuildBaseResp(err error) *knote.BaseResp {
    if err == nil {
        return baseResp(errno.Success)
    }

    e := errno.ErrNo{}
    if errors.As(err, &e) {
        return baseResp(e)
    }

    s := errno.ServiceErr.WithMessage(err.Error())
    return baseResp(s)
}

func baseResp(err errno.ErrNo) *knote.BaseResp {
    return &knote.BaseResp{
        StatusCode:    err.ErrCode,
        StatusMessage: err.ErrMsg,
    }
}
