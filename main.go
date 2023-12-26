package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("good morning %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("good bye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 { // we need to specify the return type too
	return math.Pi * r * r // uses a math library function to get the pi value
}
func main() {
	sayGreeting("nibin")
	sayBye("nibin")

	names := []string{"mario", "luigi", "bower", "peach"}
	cycleNames(names, sayGreeting) // we pass the function name only , doesnt invoke it here
	cycleNames(names, sayBye)      // we pass the function name only , doesnt invoke it here

	a1 := circleArea(10.4)
	a2 := circleArea(15)

	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f \n", a1, a2)

}
