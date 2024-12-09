package service

import (
    "context"
    "crypto/md5"
    "fmt"
    "github.com/gwen0x4c3/easy_note/cmd/user/pack"
    "io"

    "github.com/gwen0x4c3/easy_note/cmd/user/dal/mysql"
    "github.com/gwen0x4c3/easy_note/kitex_gen/kuser"
    "github.com/gwen0x4c3/easy_note/pkg/errno"
)

type UserService struct {
}

var UserServiceImpl *UserService

func NewUserService() *UserService {
    return &UserService{}
}

func init() {
    UserServiceImpl = NewUserService()
}

func (s *UserService) CreateUser(ctx context.Context, req *kuser.CreateUserRequest) error {
    user, err := mysql.QueryUser(ctx, req.UserName)
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
    return mysql.CreateUser(ctx, []*mysql.User{{
        UserName: req.UserName,
        Password: password,
    }})
}

func (s *UserService) MGetUser(ctx context.Context, req *kuser.MGetUserRequest) ([]*kuser.User, error) {
    dbUsers, err := mysql.MGetUsers(ctx, req.UserIds)
    if err != nil {
        return nil, err
    }
    res := pack.Users(dbUsers)
    return res, nil
}

func (s *UserService) CheckUser(ctx context.Context, req *kuser.CheckUserRequest) (int64, error) {
    h := md5.New()
    if _, err := io.WriteString(h, req.Password); err != nil {
        return 0, err
    }
    password := fmt.Sprintf("%x", h.Sum(nil))

    userName := req.UserName
    users, err := mysql.QueryUser(ctx, userName)
    if err != nil {
        return 0, err
    } else if len(users) == 0 {
        return 0, errno.AuthorizationFailedErr.WithMessage("用户不存在")
    }
    user := users[0]
    if user.Password != password {
        return 0, errno.AuthorizationFailedErr.WithMessage("用户名和密码不匹配")
    }
    return int64(user.ID), nil
}
