package service

import (
	"blog/dao"
	"blog/model"
	"fmt"
	"testing"
)

var dns = "root:521425Yong@tcp(127.0.0.1:3306)/blogger?parseTime=true"

func init() {
	err := dao.Init(dns)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetArticleRecordList(t *testing.T) {
	list, err := GetArticleRecordList(0, 10)
	if err != nil {
		fmt.Println(err)
	}
	for _, i := range list {
		fmt.Println(i)
	}
}

func TestGetCategoryIds(t *testing.T) {
	list, _ := dao.GetArticleList(0, 10)
	//var articleList []*model.ArticleInfo
	//articleList = append(articleList, articleInfo)
	for _, l := range list {
		fmt.Printf("%v", l)
	}
	ids := GetCategoryIds(list)
	for _, id := range ids {
		fmt.Println(id)
	}

}

func TestInsertNewArticle(t *testing.T) {
	k := &model.Category{
		CategoryId:   1,
		CategoryName: "css/html",
		CategoryNo:   1,
	}
	result, err := InsertNewArticle("AC", "BB", "CC", k)
	if err != nil {
		fmt.Println(err)
	}
	_, err = GetArticleRecordList(0, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("文章插入结果：%v\n", result)
	//for _, i := range list {
	//	fmt.Println(i)
	//}

}
