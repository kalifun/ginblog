package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Conf *ini.File

	RunMode string

	PageSize  int
	JwtSecret string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

func init() {
	var err error
	Conf, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini' :%v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Conf.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Conf.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server' :%v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Conf.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app' :%v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
