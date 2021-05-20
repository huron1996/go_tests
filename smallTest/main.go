package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "hello"
	string2 := "world"
	strs := []string{string1, string2}

	result := strings.Join(strs, "")
	fmt.Println(result)
}
