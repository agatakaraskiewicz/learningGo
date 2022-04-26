package _2shortDeclarations

import "fmt"

func shortDeclaration() (int, int) {
	x := 42
	fmt.Println(x)
	x = 23
	fmt.Println(x)

	y := 100 + 14
	return x, y
}
