package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// DSN:Data Source Name
	dsn := "root:1234@tcp(192.168.10.216:3306)/bintest"
	db, err := sql.Open("mysql", dsn) // 使用mysql数据库驱动
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err) 
	}
	// 定义结构体
	type user struct {
		class_name string
		aaa int32
	}

	// sqlStr := "select * from user where host=?"
	var u user
	// 单行查询
	rows, _ := db.Query("select class_name,aaa from test316;")
	for rows.Next() {
		rows.Scan(&u.class_name, &u.aaa)
		fmt.Printf("class_name: %s  aaa: %d \n",u.class_name, u.aaa)
	}
	if err != nil {
		fmt.Printf("scan err %v \n", err) // scan err sql: no rows in result set 
	}
}
