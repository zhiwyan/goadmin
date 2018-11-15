package mysql

import (
	"classroom/lib/config"
	"classroom/lib/mysql"
	"fmt"
	"log"
	"testing"
	"time"
)

func init() {
	configpath := "../config/classroom_settle_consumer.toml"
	err := config.InitConfig(configpath)
	if err != nil {
		panic(err)
	}
	log.Println(config.Config)

	err = mysql.InitMySQL()
	if err != nil {
		panic(err)
	}
}

func TestSaveLogInfo(t *testing.T) {
	c := new(CourseLog)
	c.CourseId = "11111"
	fmt.Println("err: ", SaveLogInfo(&CourseLog{UserId: 4, CourseId: "444", Action: 4, Content: "我爱北京天安门44", CreateTime: time.Now()}))
}
