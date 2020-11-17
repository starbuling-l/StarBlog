package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	PageSide  int
	JwtSecret string

	RunMode      string
	HTTPPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
)

func init() {
	var err error
	if Cfg, err = ini.Load("conf/app.ini"); err != nil {
		log.Fatalln("Fail to parse conf/app.ini: %v",err)
	}

	LoadBase()
	LoadApp()
	LoadServer()

}
