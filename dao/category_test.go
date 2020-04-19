package dao

import (
	"fmt"
	"testing"
)

var dns = "root:521425Yong@tcp(127.0.0.1:3306)/blogger?parseTime=true"

func TestGetCategoryAllList(t *testing.T) {
	err := Init(dns)
	if err != nil {
		panic(err)
		return
	}
	list, err := GetCategoryAllList()
	if err != nil {
		panic(err)
		return
	}
	for k, v := range list {
		fmt.Printf("%d is %#v.\n", k, v)
	}
}

func TestGetCategoryList(t *testing.T) {
	ids := []int64{1, 2}
	list, err := GetCategoryList(ids)
	if err != nil {
		fmt.Println(err)
	}
	for _, i := range list {
		fmt.Println(i)
	}
}
