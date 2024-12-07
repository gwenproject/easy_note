package main

import (
	"context"
	kuser "github.com/gwen0x4c3/easy_note/kitex_gen/kuser"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *kuser.CreateUserRequest) (resp *kuser.CreateUserResponse, err error) {
	// TODO: Your code here...
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
