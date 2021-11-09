package core

import (
	"crypto/md5"
	"fmt"
)

func Md5dd() {
	var ddf string
	fmt.Scanln(&ddf)
	data := []byte(ddf)
	s := fmt.Sprintf("%x", md5.Sum(data))
	fmt.Println(s)
	/*
		// 也可以用这种方式
		h := md5.New()
		h.Write(data)
		s = hex.EncodeToString(h.Sum(nil))
		fmt.Println(s)*/
}
