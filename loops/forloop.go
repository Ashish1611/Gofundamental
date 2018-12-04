package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		 fmt.Printf(" %d", i)

		for j := 10; j <= 20; j++ {
			fmt.Printf(" %d", j)
		}
	}
}
