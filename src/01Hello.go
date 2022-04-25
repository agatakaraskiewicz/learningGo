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
}

func foo() {
	fmt.Println("I'm in foo")
}
