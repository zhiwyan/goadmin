package filters

import (
	"config_server/lib/common"
	"config_server/lib/http"

	"github.com/gin-gonic/gin"
)

func RouterNotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// end
		http.ResponseError(c, common.ERR_NOT_FOUND_REQUEST, common.JsonEmptyObj)
		c.Abort()
	}
}
