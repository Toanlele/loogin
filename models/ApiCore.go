package models

import (
	"database/sql"
	"fmt"
)

func Apicore(INname string, Skey string) (string, string, string) {
	Runapi, err := sql.Open("mysql", Mysql)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		fmt.Println("There may be an error on the servern:")
		fmt.Println("Please check your server")
	}
	user := new(APIDATA)
	row := Runapi.QueryRow("select * from LG_Userdata where user=?", INname)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.ID, &user.Name, &user.User, &user.Hashkey, &user.Email, &user.Phone, &user.addTime); err != nil {
		fmt.Printf("scan failed, err:%v", err)
	}
	if Skey == user.Hashkey {
		return "未知", "nul", "nul"

	} else {

		return user.Name, user.Email, user.Phone
	}
}
