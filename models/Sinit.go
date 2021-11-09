//Silit文件用于验证用户名与数据库是否匹配

package models

import (
	"fmt"
)

//我定义特殊FFE数据类用于验证
/*const (
	FFEUser   string = Query()
	FFEPasswd string = "zxcvbnm"
)*/

func EffUser(Suser string, Spasswd string) string {
	var rouse string
	UserU, UserP := Query(Suser)
	if Suser == UserU {
		//验证用户是否存在,如果存在就返回参数Pass，就后继续下一步验证密码
		rouse = "Pass"
		if Spasswd == UserP {
			rouse = "Pass"
			return rouse
		} else {
			rouse = "Fail"

		}

	} else {
		rouse = "Fail"

	}
	fmt.Println(rouse)
	return rouse
}
