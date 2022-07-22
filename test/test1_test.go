package test

import (
	"fmt"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	start := time.Now()
	var sum int64
	for i := int64(0); i < 100000000; i++ {
		sum += i
	}
	d := time.Since(start)
	d2 := time.Now().Sub(start)
	fmt.Println("间隔：", d, ", sum:", sum)
	fmt.Println("间隔：", d2, ", sum:", sum)
}
