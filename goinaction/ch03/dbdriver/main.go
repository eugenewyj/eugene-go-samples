package main

import (
	"database/sql"
	_ "github.com/eugenewyj/go-sample/goinaction/ch03/dbdriver/postgres"
	"log"
)

func main() {
	db, err := sql.Open("postgres", "mydb")
	if err != nil {
		log.Println("出错", err.Error())
		return
	}
	log.Print("数据库", db)
	result, err := db.Exec("select 1 from ")
	if err != nil {
		log.Println("出错2", err.Error())
		return
	}
	log.Println("结果：", result)
}
