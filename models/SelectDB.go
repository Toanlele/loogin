package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Query(INname string) (string, string) {
	DB, err := sql.Open("mysql", Mysql)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		fmt.Println("There may be an error on the servern:")
		fmt.Println("Please check your server")
	}
	fmt.Println("1:", INname)
	AU, AP := GueryOne(DB, INname)
	//_, AP := GueryOne(DB)
	//fmt.Println("密码", GueryPasswd(DB))
	return AU, AP
	//Adduid(DB)
}

func GueryOne(DB *sql.DB, UserQU string) (string, string) {
	user := new(UserDATA)
	row := DB.QueryRow("select * from LG_Userdata where user=?", UserQU)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.ID, &user.Name, &user.User, &user.Hashkey, &user.Email, &user.Phone, &user.addTime); err != nil {
		fmt.Printf("scan failed, err:%v", err)
	}

	fmt.Println(user)
	YUuser = user.Name //调用config全局变量赋值
	YUemail = user.Email
	YUphone = user.Phone
	return user.User, user.Hashkey
}

/*func GueryList(DB *sql.DB, UserQU string) (string, string, string) {
	user := new(UserDATA)
	row := DB.QueryRow("select * from LG_Userdata where user=?", UserQU)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.ID, &user.Name, &user.User, &user.Hashkey, &user.Email, &user.Phone, &user.addTime); err != nil {
		fmt.Printf("scan failed, err:%v", err)
	}

	fmt.Println(user)
	return user.Name, user.Phone, user.Email

}*/

//Yuser,Yphone,Yemail,Yphone :=GueryList()
/*func GueryOne(DB *sql.DB) {
	user := new(User)
	row := DB.QueryRow("select * from My_LgUserDb where id=?", 2)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.ID, &user.Name, &user.User, &user.age, &user.sex, &user.class, &user.hobby, &user.addTime); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}

	fmt.Println(user.Name)
}
func GueryUser(DB *sql.DB) string {
	user := new(User)
	row := DB.QueryRow("select * from My_LgUserDb where id=?", 2)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.ID, &user.Name, &user.User, &user.age, &user.sex, &user.class, &user.hobby, &user.addTime); err != nil {
		fmt.Printf("scan failed, err:%v", err)
	}
	return user.User
}
func GueryPasswd(DB *sql.DB) string {
	user := new(User)
	row := DB.QueryRow("select * from My_LgUserDb where id=?", 2)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.ID, &user.Name, &user.User, &user.age, &user.sex, &user.class, &user.hobby, &user.addTime); err != nil {
		fmt.Printf("scan failed, err:%v", err)
	}
	return user.User
}
func Useradds() {

}
*/
