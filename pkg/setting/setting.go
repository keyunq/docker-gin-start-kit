package setting

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Title       string
	DBType      string
	Mysql       mysql
	ErrorLogDir string
	DebugLogDir string
	JwtSecret   string
}

type mysql struct {
	Ip       string
	Port     int
	UserName string
	Password string
	DbName   string
	Debug    bool
}

var (
	Cfg     Config
	RunMode string
)

func init() {

	var config Config
	if _, err := toml.DecodeFile("conf/db.toml", &config); err != nil {
		panic(err)
	}

	Cfg = config

	RunMode = "debug"

}
