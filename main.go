package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// -- we will make the following into a func getInput which is reusable for others too
	// fmt.Print("create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	name, _ := getInput("create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("created a new bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s- save item, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("whats the item name: ", reader)
		price, _ := getInput("whats the item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price should be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added - ", name, p)
		promptOptions(b)
	case "t":
		tip, _ := getInput("enter the tip amount ($):", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip should be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("added tip - ", t)
		promptOptions(b)
	case "s":
		fmt.Println("you chose to save the bill -", b)
	default:
		fmt.Println("that was not a valid option")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()

	promptOptions(myBill)
	fmt.Println(myBill)
}
