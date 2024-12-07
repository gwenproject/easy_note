package mysql

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gwen0x4c3/easy_note/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func InitMysql() {
	klog.Info("UserService: Start to init mysql, dsn: ", constants.MySQLDefaultDSN)
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	// 开启 tracing，为后续使用 Jaeger 或者 Zipkin 等做准备
	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}
	klog.Info("UserService: Init mysql success")
}
