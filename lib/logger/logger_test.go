package logger

import (
	"testing"
)

func init() {
	logConf := &LogConfig{}
	logConf.LogLevel = "debug"
	logConf.LogFile = "./out"
	logConf.IsDebug = 1
	err := InitloggerConf)
	if err != nil {
		panic(err)
	}
}

func TestLogger(t *testing.T) {
	Error("数据库错误 sql: ", "select * from xxx where id = 1")
}
