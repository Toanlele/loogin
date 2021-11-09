package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
func Postlogo(c *gin.Context) {
	// 从请求中读取文件
	User := c.PostForm("user")
	Userid := c.PostForm("passwd")
	Rstatus := EffUser(User, Userid)
	switch Rstatus {
	case "Pass":
		c.HTML(http.StatusOK, "head.html", gin.H{
			"name":  YUuser,
			"phone": YUphone,
			"email": YUemail,
		})
	case "Fail":
		c.HTML(http.StatusOK, "fail.html", nil)
	}
}
func GetAPi(c *gin.Context) {
	//使用格式 user?name=lisi&age=18
	ApiUser := c.Query("user")
	ApiKey := c.Query("key")
	name, email, phone := Apicore(ApiUser, ApiKey)

	c.JSON(200, gin.H{
		"User":  ApiUser,
		"Name":  name,
		"Email": email,
		"Phone": phone,
	})

}
