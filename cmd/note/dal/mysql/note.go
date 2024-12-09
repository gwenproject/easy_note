package mysql

import (
    "context"
    "github.com/gwen0x4c3/easy_note/pkg/constants"
    "github.com/gwen0x4c3/easy_note/pkg/errno"

    "gorm.io/gorm"
)

type Note struct {
    gorm.Model
    UserId  int64  `json:"user_id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

func (n *Note) TableName() string {
    return constants.NoteTableName
}

func CreateNote(ctx context.Context, notes []*Note) error {
    if err := DB.WithContext(ctx).Create(notes).Error; err != nil {
        return err
    }
    return nil
}

func MGetNotes(ctx context.Context, noteIds []int64) ([]*Note, error) {
    res := make([]*Note, 0)
    if len(noteIds) == 0 {
        return res, nil
    }
    if err := DB.WithContext(ctx).Where("id in ?", noteIds).Find(&res).Error; err != nil {
        return nil, err
    }
    return res, nil
}

func UpdateNote(ctx context.Context, note *Note) error {
    if note.ID == 0 {
        return errno.ParamErr
    }
    return DB.WithContext(ctx).Model(note).Updates(note).Error
}

func DeleteNote(ctx context.Context, noteId int64) error {
    if noteId == 0 {
        return errno.ParamErr
    }
    return DB.WithContext(ctx).Where("id = ?", noteId).Delete(&Note{}).Error
}

func QueryNote(ctx context.Context, userId int64, keyword *string, current, pageSize int) (res []*Note, total int64, err error) {
    conn := DB.WithContext(ctx).Model(&Note{}).Where("user_id=?", userId)

    if keyword != nil {
        conn = conn.Where("title like ?", "%"+*keyword+"%")
    }

    if err = conn.Count(&total).Error; err != nil {
        return
    }

    conn.Limit(pageSize).Offset((current - 1) * pageSize).Find(&res)
    return
}
