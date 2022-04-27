package customTypeEx

import "fmt"

//create your own type with underlying int
type ratto int

//create a variable of the new type
var wedel ratto

func customTypeEx() {
	//print the zero value
	fmt.Println(wedel)

	//print the type of the var
	fmt.Printf("%T\n", wedel)

	//assign a new value to the var
	wedel = 42

	//print the current value
	fmt.Println(wedel)
}
