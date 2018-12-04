package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n This line after for loop")
}

/*package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
}*/
