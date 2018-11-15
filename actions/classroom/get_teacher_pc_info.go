package classroom

import (
	"go_web_demo/lib/config"
	"go_web_demo/lib/http"
	"go_web_demo/lib/logger"

	"github.com/gin-gonic/gin"
)

type inputGetTeacherPcInfo struct {
	CurVersion string `json:"cur_version"`
}

func GetTeacherPcInfo(c *gin.Context) {
	var err error
	input := &inputGetTeacherPcInfo{}

	err = http.GetBodyParam(c, input)
	if err != nil {
		logger.Infof("GetTeacherPcInfo::input err:%v, input:%+v", err, input)
	}

	http.ResponseSuccess(c, config.Config.TeacherPcAppInfo)
}
