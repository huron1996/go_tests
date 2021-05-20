package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// i := 1 是短声明语句
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		//sep用来隔开输入字符串
		sep = " "
	}

	fmt.Print(s)
}
