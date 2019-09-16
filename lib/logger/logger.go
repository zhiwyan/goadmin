package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SugaredLogger *zap.SugaredLogger

type LogConfig struct {
	LogLevel string //info error debug warning
	LogFile  string
	IsDebug  int //0-not debug(output only file)  1-debug (output either file and stdout )
}

func InitLogger(conf *LogConfig) error {
	isDebug := true
	if conf.IsDebug != 1 {
		isDebug = false
	}
	err := initLogger(conf.LogFile, conf.LogLevel, isDebug)
	if err != nil {
		return err
	}
	log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
	return nil
}

func ZnTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func initLogger(lp string, lv string, isDebug bool) error {
	var js string
	if isDebug {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "console",
      "outputPaths": ["stdout", "%s"],
      "errorOutputPaths": ["stdout", "%s"]
      }`, lv, lp, lp)
	} else {
		js = fmt.Sprintf(`{
      "level": "%s",
      "encoding": "console",
      "outputPaths": ["%s"],
      "errorOutputPaths": ["%s"]
      }`, lv, lp, lp)
	}

	var cfg zap.Config
	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
		return err
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	//	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeTime = ZnTimeEncoder
	var err error
	var tlog *zap.Logger
	tlog, err = cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatal("init logger error: ", err)
		return err
	}
	SugaredLogger = tlog.Sugar()
	return nil
}

// Debug fmt.Sprintf to log a templated message.
func Debug(args ...interface{}) {
	SugaredLogger.Debug(args...)
}

// Info uses fmt.Sprintf to log a templated message.
func Info(args ...interface{}) {
	SugaredLogger.Info(args...)
}

// Warn uses fmt.Sprintf to log a templated message.
func Warn(args ...interface{}) {
	SendMonitor2DingDing(args)
	SugaredLogger.Warn(args...)
}

// Error uses fmt.Sprintf to log a templated message.
func Error(args ...interface{}) {
	SendMonitor2DingDing(args)
	SugaredLogger.Error(args...)
}

//// Fatal uses fmt.Sprintf to log a templated message.
//func Fatal(args ...interface{}) {
//	SendMonitor2DingDing(args)
//	SugaredLogger.Fatal(args...)
//}

// Debugf fmt.Sprintf to log a templated message.
func Debugf(format string, args ...interface{}) {
	SugaredLogger.Debugf(format, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(format string, args ...interface{}) {
	SugaredLogger.Infof(format, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(format string, args ...interface{}) {
	SendMonitor2DingDing(args)
	SugaredLogger.Warnf(format, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(format string, args ...interface{}) {
	SendMonitor2DingDing(args)
	SugaredLogger.Errorf(format, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message.
//func Fatalf(format string, args ...interface{}) {
//	SugaredLogger.Fatalf(format, args...)
//}

func SendMonitor2DingDing(args ...interface{}) {
	slice := make([]string, len(args))

	for i, v := range args {
		slice[i] = fmt.Sprint(v)
	}

	s := strings.Join(slice, ",")

	b := json.RawMessage(`
		{"msgtype": "text","text": {"content": "` + s + `"}}
	`)

	url := "https://oapi.dingtalk.com/robot/send?access_token="
	http.Post(url, "application/json", strings.NewReader(string(b))) //忽略dingding错误
}
