package main

import (
	"fmt"
)

func main() {

	days := 5

	switch days {

	case 1:
		fmt.Println("Sun")

	case 2:
		fmt.Println("Mon")

	case 3:
		fmt.Printf("Tue")

	case 4:
		fmt.Println("Wed")

	case 5:
		fmt.Print("Thu")

	case 6:
		fmt.Println("Fri")

	case 7:
		fmt.Println("Sat")

	}
}
