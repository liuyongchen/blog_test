package service

import (
	"blog/dao"
	"blog/model"
	"errors"
	"fmt"
	"time"
	//_ "github.com/go-sql-driver/mysql"
)

// 获取所有文章和对应分类信息
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 获取所有文章
	articleList, err := dao.GetArticleList(pageNum, pageSize)
	if err != nil {
		fmt.Println(err, "14")
		return
	}
	// 获取文章对应的分类列表
	categoryIds := GetCategoryIds(articleList)
	categoryList, err := dao.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Println(err, "21")
		return
	}
	for _, i := range categoryList {
		fmt.Println(i)
	}
	// 聚合文章和对应分类信息
	for _, article := range articleList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
			//Category:    nil,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
		//for _, k := range articleList {
		//	fmt.Println(k)
		//}
	}
	return
}

// 获取所有文章的分类ID
func GetCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
LABEL:
	for _, article := range articleInfoList {
		categoryId := article.CategoryId
		// 分类ID不重复
		for _, id := range ids {
			if id == categoryId {
				continue LABEL
			}
		}
		ids = append(ids, categoryId)
	}
	return
}

// 根据分类ID获取对应的分类信息
func GetArticleRecordListById(categoryId, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	// 获取所有文章
	articleList, err := dao.GetArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 获取文章对应的分类列表
	categoryIds := GetCategoryIds(articleList)
	categoryList, err := dao.GetCategoryList(categoryIds)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 聚合文章和对应分类信息
	for _, article := range articleList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {

				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

func InsertNewArticle(name, title, content string, category *model.Category) (result bool, err error) {
	articleInfo := &model.ArticleInfo{
		Title:      title,
		CreateTime: time.Now().Local(),
		UserName:   name,
	}
	NewArticle := &model.ArticleDetail{
		ArticleInfo: *articleInfo,
		Content:     content,
		Category:    *category,
	}
	//插入前检查是否存相同名称文章
	articleList, err := dao.GetArticleList(0, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, article := range articleList {
		if article.Title == title && article.UserName == name {
			err = errors.New("The same title was already in database. ")
			return
		}
	}
	//插入文章
	articleId, err := dao.InsertArticle(NewArticle)
	if err != nil {
		fmt.Println(err)
		return
	}
	//插入后查询是否插入成功
	article, err := dao.GetArticleDetail(articleId)
	if err != nil {
		fmt.Println(err)
		return
	}
	if article.Title == title && article.UserName == name {
		result = true
		return
	} else {
		err = errors.New("insert article fail!\n")
		return
	}
	return
}
