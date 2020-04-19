package controller

import (
	"blog/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IndexHandler(c *gin.Context) {
	articleRecordList, err := service.GetArticleRecordList(0, 10)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}

	// 用map构建gin.H
	// var data map[string]interface{}  = make(map[string]interface{},10)
	// data["category_list"] = categoryList
	for _, list := range articleRecordList {
		fmt.Println(list)
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}

//分类云的回调函数
func CategoryList(c *gin.Context) {
	//获取页面点击ID
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}
	articleRecordList, err := service.GetArticleRecordListById(int(categoryId), 0, 10)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}

	// 用map构建gin.H
	// var data map[string]interface{}  = make(map[string]interface{},10)
	// data["category_list"] = categoryList

	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"category_list": categoryList,
		"article_list":  articleRecordList,
	})

}

func GoToNewArticle(c *gin.Context) {
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}

	c.HTML(http.StatusOK, "views/post_article.html", gin.H{
		"category_list": categoryList,
	})
}

func SubmitNewArticle(c *gin.Context) {
	author := c.PostForm("author")
	title := c.PostForm("title")
	content := c.PostForm("content")
	categoryidStr := c.PostForm("category_id")
	categoryidInt, err := strconv.ParseInt(categoryidStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}
	category, err := service.GetCategoryById(categoryidInt)
	if err != nil {
		fmt.Println(err)
		c.HTML(http.StatusInternalServerError, "views/500.html", err)
	}
	result, err := service.InsertNewArticle(author, title, content, category)
	if err != nil {

		c.HTML(http.StatusInternalServerError, "views/500.html", err)
		return
	}
	if result {
		articleRecordList, err := service.GetArticleRecordList(0, 10)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "views/500.html", err)
			return
		}
		categoryList, err := service.GetAllCategoryList()
		if err != nil {
			c.HTML(http.StatusInternalServerError, "views/500.html", err)
			return
		}
		c.HTML(http.StatusFound, "views/index.html", gin.H{
			"result":        result,
			"category_list": categoryList,
			"article_list":  articleRecordList,
		})
	} else {
		categoryList, err := service.GetAllCategoryList()
		if err != nil {
			c.HTML(http.StatusInternalServerError, "views/500.html", err)
			return
		}
		c.HTML(http.StatusOK, "views/post_article.html", gin.H{
			"result":        result,
			"category_list": categoryList,
		})
	}
}
