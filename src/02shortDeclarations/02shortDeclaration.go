package _2shortDeclarations

import "fmt"

func shortDeclaration() int {
	x := 42
	fmt.Println(x)
	x = 23
	fmt.Println(x)
	return x

}
