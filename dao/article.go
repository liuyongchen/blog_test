package dao

import "blog/model"

//Id           int64     `db:"id"`
//CategoryId   int64     `db:"category_id"`
//Summary      string    `db:"summary"`
//Title        string    `db:"title"`
//ViewCount    uint32    `db:"view_count"`
//CreateTime   time.Time `db:"create_time"`
//CommentCount uint32    `db:"comment_count"`
//UserName     string    `db:"username"`
//ArticleInfo
//Content string `db:"content"`
//Category

func InsertArticle(article *model.ArticleDetail) (insertId int64, err error) {
	result, err := DB.Exec("insert into article(category_id,summary,title,view_count,comment_count,username,content)"+
		" value (?,?,?,?,?,?,?)", article.ArticleInfo.CategoryId, article.Summary, article.Title, article.ViewCount,
		article.CommentCount, article.UserName, article.Content)
	if err != nil {
		panic(err)
	}
	insertId, err = result.LastInsertId()
	//if err != nil {
	//	panic(err)
	//}
	return
}

// 获取分页所有文章，
// 当前页每页几条
func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 && pageSize < 0 {
		return
	}
	err = DB.Select(&articleList, "select id,summary,title,view_count,comment_count,username,category_id,"+
		"create_time from article where status=1 order by create_time desc limit ?,?", pageNum, pageSize)
	return
}

// 根据文章ID查文章
func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	articleDetail = new(model.ArticleDetail)
	err = DB.Get(articleDetail, "select id,summary,title,view_count,comment_count,username,category_id,create_time,"+
		"content from article where status=1 and id=?", articleId)
	return
}

// 根据分类ID查文章
func GetArticleListByCategoryId(cateoryId, pageNum, pageSize int) (articleInfoList []*model.ArticleInfo, err error) {
	if pageNum < 0 && pageSize < 0 {
		return
	}
	// articleInfoList = make([]*model.ArticleInfo, 10)
	err = DB.Select(&articleInfoList, "select id,summary,title,view_count,comment_count,username,category_id,create_time"+
		" from article where status=1 and category_id=? order by create_time desc limit ?,?", cateoryId, pageNum, pageSize)
	return
}
