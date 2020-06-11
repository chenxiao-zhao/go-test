package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	hello := "Hello "
	if len(os.Args) > 1 {
		hello += strings.Join(os.Args[1:], " ")
	}
	fmt.Println(hello)
}
