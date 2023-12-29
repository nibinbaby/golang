package main

import "fmt"

func updateName(n *string) {
	*n = "wedge"
}

func main() {

	name := "tifa"

	fmt.Println("memoy address of name is : ", &name)

	m := &name

	fmt.Println("memory address in the variable m : ", m)
	fmt.Println("value at memory address ", m, " is ", *m)

	fmt.Println(name)

	updateName(m)

	fmt.Println(name)
}

/*

|--name--|---m---|
| 0x001  | 0x002 |
|--------|-------|
| tifa   | p0x001|
|--------|-------|

*/
