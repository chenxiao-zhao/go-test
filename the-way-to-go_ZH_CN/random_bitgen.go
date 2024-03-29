package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for {
			fmt.Print(<-c, " ")
		}
	}()
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}
