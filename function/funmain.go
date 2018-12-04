package main

import (
	"fmt"
	function "function/TestFun"
)

func main() {
	price, no := 90, 6
	totalPrice := function.CalculateBill(price, no)
	fmt.Println("Total price is", totalPrice)
}
