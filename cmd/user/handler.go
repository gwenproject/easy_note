package main

import (
	"context"

	"github.com/gwen0x4c3/easy_note/cmd/user/pack"
	"github.com/gwen0x4c3/easy_note/cmd/user/service"
	kuser "github.com/gwen0x4c3/easy_note/kitex_gen/kuser"
	"github.com/gwen0x4c3/easy_note/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *kuser.CreateUserRequest) (resp *kuser.CreateUserResponse, err error) {
	resp = new(kuser.CreateUserResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}

	err = service.NewUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *kuser.MGetUserRequest) (resp *kuser.MGetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *kuser.CheckUserRequest) (resp *kuser.CheckUserResponse, err error) {
	// TODO: Your code here...
	return
}
