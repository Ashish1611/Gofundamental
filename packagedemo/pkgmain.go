package main

import (
	"fmt"
)

func main() {
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used
	 */
	fmt.Printf("area of rectangle %.2f\n", packagedemo.rectangle.Area(rectLen, rectWidth))
	/*Diagonal function of rectangle package used
	 */
	fmt.Printf("diagonal of the rectangle %.2f ", packagedemo.rectangle.Diagonal(rectLen, rectWidth))
}
