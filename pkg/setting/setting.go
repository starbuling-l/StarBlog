package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

/*var (
	Cfg *ini.File

	PageSize  int
	JwtSecret string

	RunMode string

	HTTPPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration

	Type        string
	User        string
	PassWord    string
	Host        string
	Name        string
	TablePrefix string
)

func init() {
	var err error
	if Cfg, err = ini.Load("conf/app.ini"); err != nil {
		log.Fatalf("Fail to parse conf/app.ini: %v", err)
	}

	LoadBase()
	LoadApp()
	LoadServer()
	LoadDataBase()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadApp() {
	PageSize = Cfg.Section("app").Key("PAGE_SIZE").MustInt(10)
	JwtSecret = Cfg.Section("app").Key("JWT_SECRET").MustString("23347$040412")
}

func LoadServer() {
	HTTPPort = Cfg.Section("server").Key("HTTP_PORT").MustInt(10)
	ReadTimeOut = time.Duration(Cfg.Section("server").Key("READ_TIMEOUT").MustInt(10)) * time.Second
	WriteTimeOut = time.Duration(Cfg.Section("server").Key("WRITE_TIMEOUT").MustInt(10)) * time.Second
}

func LoadDataBase() {
	Type = Cfg.Section("database").Key("TYPE").MustString("mysql")
	User = Cfg.Section("database").Key("USER").MustString("root")
	PassWord = Cfg.Section("database").Key("PASSWORD").MustString("root")
	Host = Cfg.Section("database").Key("HOST").MustString("127.0.0.1:3306")
	Name = Cfg.Section("database").Key("NAME").MustString("blog")
	TablePrefix = Cfg.Section("database").Key("TABLE_PREFIX").MustString("blog_")
}
*/

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var config *ini.File

func SetUp() {
	var err error
	config, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup fail to  parse 'conf/app.ini':%v", err)
	}
	MapTo("app", AppSetting)
	MapTo("server", ServerSetting)
	MapTo("database", DatabaseSetting)
	MapTo("redis", RedisSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	ServerSetting.ReadTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

//ini 文件必须为大写的驼峰名命法
func MapTo(section string, v interface{}) {
	if err := config.Section(section).MapTo(v); err != nil {
		log.Fatalf("config.MapTo %s err: %v", section, err)
	}
}
