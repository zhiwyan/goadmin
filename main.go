package main

import (
	"context"
	"flag"
	"goadmin/app/routers"
	"goadmin/lib/config"
	libhttp "goadmin/lib/http"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	configpath := flag.String("f", "./config/config_server.toml", "config file")
	flag.Parse()

	err := Init(*configpath)
	if err != nil {
		panic(err)
	}

	//设置模式
	gin.SetMode(config.GetMode())

	router := gin.Default()
	router.Use(libhttp.Cors())

	////session
	//store, err := sessionRedis.NewStoreWithPool(redis.RedisPool, []byte("secret"))
	//if err != nil {
	//	panic(err)
	//}
	//router.Use(sessions.Sessions("wenba_session", store))

	// 注册路由
	routers.Init(router)

	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	port := config.GetPort()

	srv := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen port %s error", port)
		}
	}()

	graceShutDown(srv)
}

/**
 * 释放连接池资源，优雅退出
 * @param  {[type]} srv *http.Server  [description]
 * @return {[type]}     [description]
 */
func graceShutDown(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	<-quit

	log.Println("Shutdown Server ...")

	//mysql.CloseMysqlPool()          //close Mysql Pool
	//redis.CloseRedisPool()          //close Redis pool
	//redis.CloseClassRoomRedisPool() // close classroom redis pool

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

/**
 * 初始化 日志，mysql，redis，NSQ
 */
func Init(configpath string) error {
	var err error

	//初始化配置文件
	err = config.InitConfig(configpath)
	if err != nil {
		return err
	}

	// initLogger()
	// defer logger.SugaredLogger.Sync()

	////初始化MySQL
	//err = mysql.InitMySQL()
	//if err != nil {
	//	return err
	//}
	//
	////初始化Redis
	//err = redis.InitRedisPool()
	//if err != nil {
	//	return err
	//}
	//
	//err = redis.InitClassRoomRedisPool()
	//if err != nil {
	//	return err
	//}
	//
	////初始化NSQ
	//err = libnsq.InitNsqProducer()
	//if err != nil {
	//	return err
	//}

	return nil
}
