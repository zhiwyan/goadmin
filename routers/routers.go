package routers

import (
	"config_server/actions/classroom"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine) (err error) {
	if g == nil {
		err = fmt.Errorf("nil gin engine")
		return
	}

	//rg := g.Group("test")
	//rg.Use(filters.RouterNotFound())

	g.POST("/teacherPcAppInfo/getTeacherPcInfo", classroom.GetTeacherPcInfo)
	g.POST("/studentPcAppInfo/getStudentPcInfo", classroom.GetStudentPcInfo)

	return
}
