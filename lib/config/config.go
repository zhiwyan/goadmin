package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

var Config Configure

type Configure struct {
	Env string
	//Mysql          MysqlConfig
	//Redis          RedisConfig
	//ClassRoomRedis RedisConfig
	Mode   string
	Port   int
	AppKey string

	//NsqAddr        string
	//NsqLookupdAddr string

	Log                   LogConf
	TeacherPcAppInfo      TeacherPcAppInfoConfig
	StudentPcAppInfo      StudentPcAppInfoConfig
	StudentAndroidInfo    StudentAndroidIosAppInfoConfig
	StudentAndroidAppInfo StudentAndroidIosAppInfoConfig
	StudentIosInfo        StudentAndroidIosAppInfoConfig
	StudentIosAppInfo     StudentAndroidIosAppInfoConfig
}

type TeacherPcAppInfoConfig struct {
	AppUrl       string
	ForcedUpdate int
	LastVersion  string
	Md5          string
	UpdateInfo   string
	New          int
	ForgetPage   string
	SignupPage   string
}

type StudentPcAppInfoConfig struct {
	AppUrl       string
	ForcedUpdate int
	LastVersion  string
	Md5          string
	UpdateInfo   string
	New          int
	ForgetPage   string
	SignupPage   string
}

type StudentAndroidIosAppInfoConfig struct {
	AppID        string
	AppURL       string
	ForcedUpdate int
	LastVersion  string
	Md5          string
	MinVersion   string
	New          int
	Size         string
	UpdateInfo   string
}

type MysqlConfig struct {
	Host         string
	Password     string
	User         string
	Database     string
	Maxlifetime  int
	MaxidleConns int
}

type RedisConfig struct {
	Server              string
	Password            string
	RedisMaxIidleConn   int
	RedisIdleTimeoutSec int
	SelectDb            int
}

type LogConf struct {
	LogLevel string
	LogFile  string
	IsDebug  int
}

func InitConfig(configpath string) error {
	configBytes, err := ioutil.ReadFile(configpath)
	if err != nil {
		return err
	}

	if _, err := toml.Decode(string(configBytes), &Config); err != nil {
		return err
	}
	log.Println(Config)

	return nil
}

func GetMode() string {
	return Config.Mode
}

func GetPort() string {
	return fmt.Sprintf(":%d", Config.Port)
}
