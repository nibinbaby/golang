package main

import "fmt"

func main() {

	myBill := newBill("marios bill")

	myBill.updateTip(10)

	myBill.addItem("muffin", 3.56)

	fmt.Println(myBill.format())

}
