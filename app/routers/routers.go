package routers

import (
	"goadmin/app/controller"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	// 设置静态资源路由
	InitStaticRouter(router)

	// 设置小工具路由
	r := router.Group("/")
	{

		r.GET("/", controller.Index.Index)
		r.GET("/main", controller.Index.Main)

		// 菜单
		// menu := r.Group("/menu")
		// {
		// 	menu.GET("/", controller.Menu.Index)
		// 	menu.GET("/index", controller.Menu.Index)
		// 	menu.POST("/list", controller.Menu.List)
		// }
	}

}

func InitStaticRouter(router *gin.Engine) {
	// 设置静态资源路由
	router.Static("/resource", "./public/resource")
	router.StaticFile("/favicon.ico", "./public/resource/images/favicon.ico")
	router.LoadHTMLGlob("./public/views/*")
}
