package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 5; i++ {
		 fmt.Printf(" %d", i)

		for j := 5; j <= 10; j++ {
			fmt.Printf(" %d", j)
		}
	}
}
