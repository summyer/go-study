package util2_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestStrings(t *testing.T) {
	var s = " this is demo "
	fmt.Println(s)
	s = strings.Trim(s, " ")
	fmt.Println(s)
	fmt.Println(strings.ToTitle(s))
}

func TestString2(t *testing.T) {
	var s strings.Builder
	s.WriteString("hello")
	s.WriteString("world")
	fmt.Println(s.String())
}
