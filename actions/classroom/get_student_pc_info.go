package classroom

import (
	"config_server/lib/config"
	"config_server/lib/http"
	"config_server/lib/logger"

	"github.com/gin-gonic/gin"
)

type inputGetStudentPcInfo struct {
	CurVersion string `json:"cur_version"`
}

func GetStudentPcInfo(c *gin.Context) {
	var err error
	input := &inputGetStudentPcInfo{}

	err = http.GetBodyParam(c, input)
	if err != nil {
		logger.Infof("GetStudentPcInfo::input err:%v, input:%+v", err, input)
	}

	http.ResponseSuccess(c, config.Config.StudentPcAppInfo, input)
}
