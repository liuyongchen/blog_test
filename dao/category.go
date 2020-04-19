package dao

import (
	"blog/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//获取单个分类
func GetCategoryById(id int64) (category *model.Category, err error) {
	DB.Get(&category, "select id,category_name,category_no from category where id=?", id)
	return
}

//获取多个分类
func GetCategoryList(categoryIds []int64) (categoryList []*model.Category, err error) {
	sqlStr, args, err := sqlx.In("select id,category_name,category_no from category where id in(?)", categoryIds)
	err = DB.Select(&categoryList, sqlStr, args...)
	return
}

func GetCategoryAllList() (categoryList []*model.Category, err error) {
	DB.Select(&categoryList, "select id,category_name,category_no from category")
	return
}
