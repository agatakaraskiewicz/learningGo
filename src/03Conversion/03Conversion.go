package _3Conversion

import "fmt"

var milky int

type rat int

var oreo rat

func conversion() (int, int, int) {
	milky = 2
	oreo = 5

	knoppers := int(oreo)
	fmt.Println(knoppers)

	x, _ := fmt.Printf("%T\n", milky)
	fmt.Println(x)

	return x, milky, knoppers

}
