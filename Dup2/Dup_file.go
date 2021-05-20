package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}
}

func main() {
	counts := make(map[string]int)
	//	命令行的参数
	file := os.Args[1:]
	//如果命令没有参数，就等待Stdin输入
	if len(file) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range file {
			//读取arg
			f, err := os.Open(arg)

			//如果出错
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
