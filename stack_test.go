package mygoutil

import (
	"fmt"
	"testing"
)

func TestStackAll(t *testing.T) {
	s := StackInt{}
	for i := 0; i < 10; i++ {
		s.push(i)
	}
	for !s.empty() {
		fmt.Println(s.top())
		s.pop()
	}
}