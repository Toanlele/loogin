package models

import (
	"github.com/gin-gonic/gin"
)

func Corenmdd() { //
	r := gin.Default()
	r.Static("/src", "src")
	//r.LoadHTMLGlob("./templates/*")
	r.LoadHTMLGlob("src/**/*")
	r.GET("/admin", Admin)
	r.POST("/home", Postlogo)
	r.GET("/api", GetAPi)
	r.Run("0.0.0.0:8888")
}
