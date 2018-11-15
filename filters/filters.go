package filters

import (
	"config_server/lib/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 设置新旧版本cookie
func SetServerVersion() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := &http.Cookie{
			Name:  "server_version",
			Value: "1.0",
		}
		http.SetCookie(c.Writer, cookie)
		c.Next()
	}
}

// 设置程序开始时间
func SetRequesetTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("reqStartTime", common.Start())
		c.Next()
	}
}
