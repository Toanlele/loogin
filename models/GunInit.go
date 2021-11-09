package models

import "fmt"

const (
	YesINit int = 1
	NoInit  int = 0
)

func GInit(status int) string {
	var ss string
	if status != YesINit {
		fmt.Println("") //判断程序是否需要初始化

	}
	if status != NoInit {
		fmt.Println("不")
		ss = "沙比"
	}
	return ss
}
