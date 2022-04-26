package main

import "fmt"

func main() {
	fmt.Println("Hello Worldy")

	foo()

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	//.Println returns n (number of bytes in the printed msg) and err (error if any)
	n, err := fmt.Println("Hi")
	fmt.Println(n)
	fmt.Println(err)
}

func foo() {
	fmt.Println("I'm in foo")
}
