package main

import "fmt"

func main() {
	age := 45
	name := "Someone"
	// Print - it starts with capital P because it is a public function from the fmt package
	fmt.Print("hello, ")
	fmt.Print("world ! \n")
	fmt.Print("new line \n")

	// a new line character is automatically appended to the printed line
	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas!")

	fmt.Println("my name is ", name, "and age is", age)

	// Printf (formatted strings) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name) // quotes around the value
	fmt.Printf("age is of type is %T \n", age)                 // give the type of the variable value
	fmt.Printf("a floating num is %f \n", 4.756567)
	fmt.Printf("a floating num is %0.1f \n", 4.756567) // round the number to the points

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)

	fmt.Println("saved string to variable str is : ", str)
}
