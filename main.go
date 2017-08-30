package main

import (
	. "github.com/leveldorado/linear/vector"
	"fmt"
)

func main() {
	fmt.Println(Vec{-0.221, 7.437}.Magnitude())
	fmt.Println(Vec{8.813, -1.331, -6.247}.Magnitude())
	fmt.Println(Vec{5.581, -2.136}.Normalize())
	fmt.Println(Vec{1.996, 3.108, -4.554}.Normalize())
}
