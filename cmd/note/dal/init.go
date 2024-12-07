package dal

import (
	"github.com/gwen0x4c3/easy_note/cmd/note/dal/mysql"
)

func init() {
	mysql.InitMysql()
}
