package varDeclarationEx

import "fmt"

//declare 3 vars with 3 different types ad package lvl scope
var age int
var person string
var bald bool

//print the default values assigned by compiler -> zero values
func varDeclarationEx() {
	fmt.Println(age)
	fmt.Println(person)
	fmt.Println(bald)

	//Assign values to the variables
	age = 42
	person = "James Bond"
	bald = true

	//print all of the values to a single string
	x := fmt.Sprintf("%v, %v, %v", age, person, bald)
	fmt.Println(x)

	//print the type of x to be sure that it is a string
	fmt.Printf("%T\n", x)
}
