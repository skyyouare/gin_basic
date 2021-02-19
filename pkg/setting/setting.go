package setting

import (
	"github.com/BurntSushi/toml"
	"log"
	"time"
)

var (
	cfg           = &Config{}
	AppSetting    = &App{}
	ServerSetting = &Server{}
	MysqlSetting  = &Mysql{}
	RedisSetting  = &Redis{}
)

//conf
type Config struct {
	App    *App
	Server *Server
	Mysql  *Mysql
	Redis  *Redis
}

//app
type App struct {
	DebugMode    string
	TimeLocation string
}

//server
type Server struct {
	HttpPort string
}

//mysql
type Mysql struct {
	UserName        string
	PassWord        string
	IpHost          string
	Port            string
	DbName          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

//redis
type Redis struct {
	IpHost string
}

//setup conf
func Setup(confPath string) {
	if _, err := toml.DecodeFile(confPath, &cfg); err != nil {
		log.Fatal(err)
	}
	AppSetting = cfg.App
	ServerSetting = cfg.Server
	MysqlSetting = cfg.Mysql
	RedisSetting = cfg.Redis
	return
}
