package routers

import (
	"fmt"
	"go_web_demo/actions/classroom"
	"go_web_demo/filters"

	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine) (err error) {
	if g == nil {
		err = fmt.Errorf("nil gin engine")
		return
	}

	//rg := g.Group("test")
	g.Use(filters.SetRequesetTime())

	g.POST("/teacherPcAppInfo/getTeacherPcInfo", classroom.GetTeacherPcInfo)
	g.POST("/studentPcAppInfo/getStudentPcInfo", classroom.GetStudentPcInfo)

	return
}
