package main

import (
	"Ashish/Gofundamental/Addition"
	"TestPackage/Division"
	"TestPackage/Modus"
	"TestPackage/Multiplication"
	"TestPackage/Substarction"
	"fmt"
)

func main() {

	var a float64
	var b float64

	choice := 0

	fmt.Println("\nWhich Operation Do you want to perform")

	fmt.Println("1.Addition", "\n2.Substarction", "\n3.Multiplication", "\n4.Division", "\n5.Modus")

	fmt.Scanf("%d", &choice)

	switch choice {

	case 1:
		fmt.Printf("Enter two No's For Addition\n")

		fmt.Scanf("%f %f", &a, &b)

		fmt.Printf("\nAddition is", Addition.Add(a, b))

	case 2:
		fmt.Printf("Enter two No's For Substraction\n")
		fmt.Scanf("%f %f", &a, &b)
		fmt.Printf("\nSubstraction is", Substarction.Sub(a, b))

	case 3:
		fmt.Printf("Enter two No's For Multiplication\n")
		fmt.Scanf("%f %f", &a, &b)
		fmt.Printf("\nMultiplication is", Multiplication.Mul(a, b))

	case 4:
		fmt.Printf("Enter two No's For Division\n")
		fmt.Scanf("%f %f", &a, &b)
		fmt.Printf("\nDivision is", Division.Div(a, b))

	case 5:
		fmt.Printf("Enter two No's For Modus\n")
		fmt.Scanf("%f %f", &a, &b)
		fmt.Printf("\nModulas is", Modus.Modulas(a, b))

	}

}
