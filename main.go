package main

import "fmt"

// package scope

var score = 96.4 // you should give this outside the "main()" since it wont be globally available otherwise

func main() {

	// var score = 98.2 // wont work , will give undefined error while referencing in the other file
	sayHello("mari")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}

// the above code will emit erro if "go run main.go" is called
// because the "greetings.go" is never executed for the compiler to read
// so you need to run "go run main.go greetings.go"
