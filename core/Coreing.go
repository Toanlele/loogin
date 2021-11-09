package core

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Msqldd() {
	db, err := sql.Open("mysql", Mysql)
	if err != nil {
		log.Fatalln(err)
	}
	result, err := db.Exec("insert into doctor_tb(name,age,addTime) values(?,?,Now())", "熊娇", 11)
	if err != nil {
		fmt.Println("新增数据错误", err)
		return
	}
	newID, _ := result.LastInsertId() //新增数据的ID
	i, _ := result.RowsAffected()     //受影响行数
	fmt.Printf("新增的数据ID：%d , 受影响行数：%d \n", newID, i)
}
