package main

import (
	"fmt"
)

func main() {

	num := 9

	if num <= 50 {

		fmt.Print("No is Less than or equal to 50 ")
	} else if num >= 51 && num <= 100 {
		fmt.Print("No is between 51 and 100")
	} else {
		fmt.Print("No is greater than 100")
	}

}
