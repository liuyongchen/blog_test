package main

import (
	"blog/controller"
	"blog/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dns = "root:521425Yong@tcp(127.0.0.1:3306)/blogger?parseTime=true"
)

func main() {
	router := gin.Default()
	err := dao.Init(dns)
	if err != nil {
		fmt.Println(err)
		return
	}
	//加载资源
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("./views/*")
	//点击主页
	router.GET("/", controller.IndexHandler)
	//点击分类
	router.GET("/category/", controller.CategoryList)
	//点击投稿
	router.GET("/article/new/", controller.GoToNewArticle)
	router.POST("/article/submit/", controller.SubmitNewArticle)

	_ = router.Run(":8000")
}
