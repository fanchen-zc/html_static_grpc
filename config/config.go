package config

import (
	"github.com/go-ini/ini"
	"html_static_grpc/helper"
	"runtime"
)

type SysConfig struct {
	Env   string `ini:"env"`
	Debug bool   `ini:"debug"`

	DBDriver string `ini:"db_driver"`
	DBHost   string `ini:"db_host"`
	DBPort   string `ini:"db_port"`
	DBUser   string `ini:"db_user"`
	DBPass   string `ini:"db_pass"`
	DBName   string `ini:"db_name"`
	DBDebug  bool   `ini:"db_debug"`

	RedisHost string `ini:"redis_host"`
	RedisPwd  string `ini:"redis_pwd"`
	RedisDb   int    `ini:"redis_db"`

	HttpListenPort string `ini:"http_listen_port"`
	GrpcPort       string `ini:"grpc_port"`

	Domain    string `ini:"domain"`
	NotifyUrl string `ini:"notify_url"`
}

var Configs *SysConfig = &SysConfig{}

// 加载系统配置文件
func Default() {
	appDir := helper.GetAppDir()
	if runtime.GOOS != "windows" {
		appDir = helper.GetCurrentPath()
	}
	conf, err := ini.Load(appDir + "/config.ini") //加载配置文件
	if err != nil {
		panic(any(err))
	}
	conf.BlockMode = false
	err = conf.MapTo(&Configs) //解析成结构体
	if err != nil {
		panic(any(err))
	}
}
