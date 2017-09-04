package main

import (
	. "github.com/leveldorado/linear/vector"
	"fmt"
)

func main() {
	fmt.Println(Vec{-7.579, -7.88}.IsParalel(Vec{22.737, 23.64}))
	fmt.Println(Vec{-7.579, -7.88}.IsOrthogonal(Vec{22.737, 23.64}))
	fmt.Println(Vec{-2.029, 9.97, 4.172}.IsParalel(Vec{-9.231, -6.639, -7.245}))
	fmt.Println(Vec{-2.029, 9.97, 4.172}.IsOrthogonal(Vec{-9.231, -6.639, -7.245}))
	fmt.Println(Vec{-2.328, -7.284, -1.214}.IsParalel(Vec{-1.821, 1.072, -2.94}))
	fmt.Println(Vec{-2.328, -7.284, -1.214}.IsOrthogonal(Vec{-1.821, 1.072, -2.94}))
}
