package core

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Useradd() {
	db, err := sql.Open("mysql", Mysql)
	if err != nil {
		log.Fatalln(err)
	}
	//result, err := db.Exec("insert into My_LgUserDb(name,user,age,sex,class, hobby,addTime) values("张飞","xiaoming","18","男","机电一体化1班","打球",Now()))
	result, err := db.Exec("insert into My_LgUserDb(name,user,age,sex,class, hobby,addTime) values(,Now()))
	if err != nil {
		fmt.Println("新增数据错误", err)
		return
	}
	newID, _ := result.LastInsertId() //新增数据的ID
	i, _ := result.RowsAffected()     //受影响行数
	fmt.Printf("新增的数据ID：%d , 受影响行数：%d \n", newID, i)
}
func queryOne(DB *sql.DB) {
	user := new(User)
	row := DB.QueryRow("select * from users where id=?", 1)
	if err := row.Scan(&user.ID, &user.Name, &user.age, &user.class, &user.hobby); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}
	fmt.Println(*user)
}