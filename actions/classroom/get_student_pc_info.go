package classroom

import (
	"go_web_demo/lib/config"
	"go_web_demo/lib/http"
	"go_web_demo/lib/logger"

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

	http.ResponseSuccess(c, config.Config.StudentPcAppInfo)
}
