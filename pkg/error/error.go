package error

import (
	"errors"
	"fmt"
)

const (
	SuccessCode                = 0
	ServiceErrCode             = 10001
	ParamErrCode               = 10002
	UserAlreadyExistErrCode    = 10003
	AuthorizationFailedErrCode = 10004
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(errCode int64, errMsg string) ErrNo {
	return ErrNo{
		ErrCode: errCode,
		ErrMsg:  errMsg,
	}
}

func (e ErrNo) WithMessage(errMsg string) ErrNo {
	e.ErrMsg = errMsg
	return e
}

var (
	Success                = NewErrNo(SuccessCode, "Success")
	ServiceErr             = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr               = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	AuthorizationFailedErr = NewErrNo(AuthorizationFailedErrCode, "Authorization failed")
)

func ConvertErr(err error) ErrNo {
	errNo := ErrNo{}
	if errors.As(err, &errNo) {
		return errNo
	}
	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
