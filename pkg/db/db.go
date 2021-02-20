package db

import (
	"fmt"
	"gin_basic/pkg/logger"
	"gin_basic/pkg/setting"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Conn *sqlx.DB

//Setup 数据库设置
func Setup() {
	//dsn "用户名:密码@[连接方式](主机名:端口号)/数据库名"
	//var dsn = "test:ceshi@(192.168.8.33)/socialtouch_dashboard_v4_20201127_140809"
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", setting.MysqlSetting.UserName, setting.MysqlSetting.PassWord, setting.MysqlSetting.IPHost, setting.MysqlSetting.Port, setting.MysqlSetting.DbName)
	dataSourceName = dataSourceName + "?parseTime=true&loc=Asia%2FShanghai&charset=utf8"
	database, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		logger.Fatal(err)
	}
	database.SetMaxIdleConns(setting.MysqlSetting.MaxIdleConns)                     // 设置空闲链接
	database.SetMaxOpenConns(setting.MysqlSetting.MaxOpenConns)                     // 最大连接数 0为不限制
	database.SetConnMaxLifetime(setting.MysqlSetting.ConnMaxLifetime * time.Minute) //可重用链接得最大时间长度
	Conn = database
}
