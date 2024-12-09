package pack

import (
    "github.com/gwen0x4c3/easy_note/cmd/user/dal/mysql"
    "github.com/gwen0x4c3/easy_note/kitex_gen/kuser"
)

func Users(users []*mysql.User) []*kuser.User {
    res := make([]*kuser.User, 0, len(users))
    for _, u := range users {
        res = append(res, User(u))
    }
    return res
}

func User(user *mysql.User) *kuser.User {
    return &kuser.User{
        UserId:   int64(user.ID),
        UserName: user.UserName,
        Avatar:   user.Avatar,
    }
}
