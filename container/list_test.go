package container_test

import (
	"container/list"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	var l list.List
	fmt.Println(l.Len())
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	l.PushBack("d")
	fmt.Println(l.Len())
	i := l.Front()
	fmt.Println(i.Value.(string))
	i = i.Next()
	fmt.Println(i.Value.(string))

	for item := l.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value.(string))
	}

}
func TestList2(t *testing.T) {
	var l *list.List = list.New()
	fmt.Println(l.Len())
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	l.PushBack("d")
	fmt.Println(l.Len())
	i := l.Front()
	fmt.Println(i.Value.(string))
	i = i.Next()
	fmt.Println(i.Value.(string))

	for item := l.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value.(string))
	}

}
