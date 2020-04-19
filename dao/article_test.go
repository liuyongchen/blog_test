package dao

import (
	"fmt"
	"testing"
)

func TestGetArticleList(t *testing.T) {
	var dns = "root:521425Yong@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		fmt.Println(err)
	}
	list, err := GetArticleList(0, 10)
	if err != nil {
		fmt.Println(err)
	}
	for _, i := range list {
		fmt.Println(i, '\n')
	}
}

func TestGetArticleDetail(t *testing.T) {
	var dns = "root:521425Yong@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err != nil {
		fmt.Println(err)
	}
	detail, err := GetArticleDetail(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(detail)
}
