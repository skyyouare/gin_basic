package setting

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

var (
	cfg = &Config{}
	//AppSetting app配置
	AppSetting = &App{}
	//ServerSetting server配置
	ServerSetting = &Server{}
	//LogSetting log配置
	LogSetting = &Log{}
	//MysqlSetting mysql配置
	MysqlSetting = &Mysql{}
	//RedisSetting redis配置
	RedisSetting = &Redis{}
)

// Config 配置
type Config struct {
	App    *App
	Server *Server
	Log    *Log
	Mysql  *Mysql
	Redis  *Redis
}

// App app相关配置
type App struct {
	DebugMode    string
	TimeLocation string
}

// Server server相关配置
type Server struct {
	HTTPPort string
}

// Log log相关配置
type Log struct {
	FileName  string
	MaxSize   int
	LocalTime bool
	Compress  bool
}

// Mysql mysql相关配置
type Mysql struct {
	UserName        string
	PassWord        string
	IPHost          string
	Port            string
	DbName          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

// Redis redis相关配置
type Redis struct {
	IPHost string
}

//Setup 设置配置
func Setup(confPath string) {
	if _, err := toml.DecodeFile(confPath, &cfg); err != nil {
		log.Fatal(err)
	}
	AppSetting = cfg.App
	ServerSetting = cfg.Server
	LogSetting = cfg.Log
	MysqlSetting = cfg.Mysql
	RedisSetting = cfg.Redis
	return
}
