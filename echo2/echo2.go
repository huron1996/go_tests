package main

import "fmt"
import "os"

func main() {
	s, sep := "", ""

	//range返回key-value，只需要value，所以用空标识符_去接受index
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Print(s)
}
