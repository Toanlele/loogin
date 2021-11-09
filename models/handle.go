package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Crddtedb() {
	db, err := sql.Open("mysql", Mysql)
	if err != nil {
		log.Fatalln(err)
	}
	db.Exec("")
	defer db.Close()

}

func TouchInit() {
	db, err := sql.Open("mysql", Mysql) //声明数据库类型为mysql 使用用户名,密码,连接形式,链接地址和端口,指定数据库
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("数据库连接成功")
	//定义一个创建数据库表的变量,后续进行exec的调用.
	createTable := "CREATE TABLE `Lg_user` (`uid` int(10) NOT NULL AUTO_INCREMENT,`username` varchar(64) NOT NULL DEFAULT '1', `gender` tinyint(1) DEFAULT NULL,`password` varchar(64) DEFAULT NULL,`created` date DEFAULT NULL,PRIMARY KEY (`uid`)) ;"
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("数据库创建成功")
}
