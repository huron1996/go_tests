package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("This is a test\n I do not know the result exactly!\n I am huron")
	scanner := bufio.NewScanner(reader)

	fmt.Printf("This is %c\n", 10)
	fmt.Printf("This is %d\n", 10)
	fmt.Printf("This is %X\n", 10)
	fmt.Printf("This is %v\n", os.Stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
