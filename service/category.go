package service

import (
	"blog/dao"
	"blog/model"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = dao.GetCategoryAllList()
	if err != nil {
		return
	}
	return
}

func GetCategoryById(id int64) (category *model.Category, err error) {
	category, err = dao.GetCategoryById(id)
	//categoryList, err = byId, err
	if err != nil {
		return
	}
	return
}
