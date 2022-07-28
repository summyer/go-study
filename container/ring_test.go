package container_test

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	var r ring.Ring
	fmt.Println(r.Len())
	var r2 *ring.Ring = ring.New(10)
	fmt.Println(r2.Len())
	for i := 0; i < r2.Len(); i++ {
		r2.Value = i + 1
		r2 = r2.Next()
	}
	//因为是一个环，所以又回到了开始
	for i := 0; i < r2.Len(); i++ {
		fmt.Println(r2.Value)
		r2 = r2.Next()
	}
	r2.Do(func(i interface{}) {
		fmt.Println("循环：", i.(int))
	})
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println(b.N)
	}
}
