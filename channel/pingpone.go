package main

import (
	"fmt"
)

var c1 chan string

func main() {
	c1 = make(chan string)
	go PingPone()
	for i := 0; i < 10; i++ {
		c1 <- fmt.Sprintf("From main: Hello,#%d", i)
		fmt.Println(<-c1)
	}
}

func PingPone() {
	i := 0
	for {
		fmt.Println(<-c1)
		c1 <- fmt.Sprintf("From Pingpong:Hi,#%d", i)
		i++
	}
}
