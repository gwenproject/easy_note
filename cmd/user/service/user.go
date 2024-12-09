package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/gwen0x4c3/easy_note/cmd/user/dal/mysql"
	"github.com/gwen0x4c3/easy_note/kitex_gen/kuser"
	"github.com/gwen0x4c3/easy_note/pkg/errno"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{ctx: ctx}
}

func (s *UserService) CreateUser(req *kuser.CreateUserRequest) error {
	user, err := mysql.QueryUser(s.ctx, req.UserName)
	if err != nil {
		return err
	}
	if len(user) != 0 {
		return errno.UserAlreadyExistErr
	}
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	return mysql.CreateUser(s.ctx, []*mysql.User{{
		UserName: req.UserName,
		Password: password,
	}})
}
