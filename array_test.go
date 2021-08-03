package mygoutil

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestLeftFirstGreater(t *testing.T) {
	a := make([]int, 0)
	for i := 0; i < 20; i++ {
		a = append(a, rand.Intn(10))
	}
	p := LeftFirstGreater(a)
	b := VOP(a, p)
	for k := range a {
		fmt.Printf("%v\t", k)
	}
	fmt.Println()
	for _, v := range a {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()
	for _, v := range p {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()
	for _, v := range b {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()
}