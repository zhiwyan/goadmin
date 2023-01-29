package mysql

// 参考 http://go-database-sql.org/connection-pool.html
import (
	"fmt"
	"goadmin/lib/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var MysqlPool *gorm.DB

func InitMySQL() error {
	var err error
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.Config.Mysql.User, config.Config.Mysql.Password, config.Config.Mysql.Host, config.Config.Mysql.Database)

	log.Println("连接mysql: ", dns)

	MysqlPool, err = gorm.Open("mysql", dns)
	if err != nil {
		return err
	}

	//在进程退出时释放mysql连接池  在入口处，应用程序退出时
	//defer MysqlPool.DB().Close()

	// mysql Pool 根据系统状况自行配置
	//MysqlPool.DB().SetConnMaxLifetime(time.Duration(mysqlSection.Key("maxlifetime").MustInt(30)) * time.Second)

	//MysqlPool.DB().SetMaxIdleConns(mysqlSection.Key("maxidleConns").MustInt(10))

	//MysqlPool.SetMaxOpenConns(n)  //不限制数据库最大并发数

	err = MysqlPool.DB().Ping()
	if err != nil {
		return err
	}

	return nil

}

func CloseMysqlPool() {
	if MysqlPool == nil {
		return
	}
	if MysqlPool.DB() != nil {
		MysqlPool.DB().Close()
	}
}

// func GetDB() *sql.DB {
// 	return MysqlPool.DB()
// }
