package mysql

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	UserId  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (n *Note) TableName() string {
	return "note"
}
