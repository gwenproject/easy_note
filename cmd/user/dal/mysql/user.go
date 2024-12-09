package mysql

import (
    "context"

    "github.com/gwen0x4c3/easy_note/pkg/constants"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    UserName string `json:"user_name"`
    Password string `json:"password"`
    Avatar   string `json:"avatar"`
}

func (u *User) TableName() string {
    return constants.UserTableName
}

func CreateUser(ctx context.Context, users []*User) error {
    return DB.WithContext(ctx).Create(users).Error
}

func MGetUsers(ctx context.Context, userIds []int64) ([]*User, error) {
    res := make([]*User, 0)
    if len(userIds) == 0 {
        return res, nil
    }
    if err := DB.WithContext(ctx).Where("id in ?", userIds).Find(&res).Error; err != nil {
        return nil, err
    }
    return res, nil
}

func QueryUser(ctx context.Context, userName string) ([]*User, error) {
    res := make([]*User, 0)
    if err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error; err != nil {
        return nil, err
    }
    return res, nil
}
