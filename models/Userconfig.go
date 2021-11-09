package models

type Userdb struct {
	User   string
	passwd string
	secode int
}

/*type process struct {
	cmd *exec.Cmd
}
*/
const (
	Mysql = "root:你的密码@tcp(0.0.0.0:3306)/Lg_user?parseTime=true"
)

var YUuser string //用户名
var YUemail string
var YUphone string
