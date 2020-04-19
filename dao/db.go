package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
	//dns = "root@521425Yong"
)

func Init(dns string) (err error) {
	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
