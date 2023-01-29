package controller

import (
	"github.com/gin-gonic/gin"
)

var Index = new(IndexController)

type IndexController struct {
}

func (ctl *IndexController) Index(c *gin.Context) {
	c.HTML(200, "index", gin.H{
		//"userInfo": userInfo,
		// "menuList": menuList,
	})
}

func (ctl *IndexController) Main(c *gin.Context) {
	c.HTML(200, "index", gin.H{})
}
