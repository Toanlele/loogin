package models

import "fmt"

func Menu() {
	dd := 11
	fmt.Println(dd)
	ptp := Userdb{
		User:   "sss",
		passwd: "SSSS",
		secode: 11,
	}
	fmt.Println(ptp)

}

/*type User struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	User    string `db:"user"`
	age     int    `db:"age"`
	sex     string `db:"sex"`
	class   string `db:"class"`
	hobby   string `db:"hobby"`
	addTime string `db:"addTime"`
}*/
type UserDATA struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	User    string `db:"user"`
	Hashkey string `db:"hashkey"`
	Email   string `db:"email"`
	Phone   string `db:"phone"`
	addTime string `db:"addTime"`
}
type APIDATA struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	User    string `db:"user"`
	Hashkey string `db:"hashkey"`
	Email   string `db:"email"`
	Phone   string `db:"phone"`
	addTime string `db:"addTime"`
}
