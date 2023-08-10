package gorm

//gorm v2
import (
	"fmt"
	"gin_basic/pkg/logger"
	"gin_basic/pkg/setting"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Conn *gorm.DB

type Writer struct {
}

func (w Writer) Printf(format string, v ...interface{}) {
	logger.Infow("gin_gorm",
		"sql", v[3],
		"src", v[0],
		"duration", v[1],
		"values", "",
		"rows_returned", v[2],
	)
}

// Setup initializes the database instance
func Setup() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.MysqlSetting.UserName,
		setting.MysqlSetting.PassWord,
		setting.MysqlSetting.IPHost,
		setting.MysqlSetting.Port,
		setting.MysqlSetting.DbName)

	GormLogger := glogger.New(Writer{}, glogger.Config{
		SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
		LogLevel:                  glogger.Info,           // 日志级别
		IgnoreRecordNotFoundError: false,                  // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  false,                  // 禁用彩色
	})

	database, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		Logger:                 GormLogger,
		SkipDefaultTransaction: false, //跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   setting.MysqlSetting.TablePrefix, // 表名前缀，`Article` 的表名应该是 `it_articles`
			SingularTable: true,                             // 表不加s
		},
	})

	if err != nil {
		logger.Fatal("gorm.Setup err: %v", err)
	}

	database.Debug()

	sqlDB, _ := database.DB()
	sqlDB.SetMaxIdleConns(setting.MysqlSetting.MaxIdleConns)                     //设置空闲链接
	sqlDB.SetMaxOpenConns(setting.MysqlSetting.MaxOpenConns)                     //最大打开连接数 0为不限制
	sqlDB.SetConnMaxLifetime(setting.MysqlSetting.ConnMaxLifetime * time.Minute) //可重用链接得最大时间长度

	Conn = database
}

// Close closes database connection (unnecessary)
func Close() {
	sqlDB, _ := Conn.DB()
	defer sqlDB.Close()
}
