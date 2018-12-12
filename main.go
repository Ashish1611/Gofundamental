package main

import (
	"Ashish/Gofundamental/Addition"
	"Ashish/Gofundamental/function/TestFun"
	"Ashish/Gofundamental/function/multivalue"
	"Ashish/Gofundamental/ifstmt/elseif"
	"Ashish/Gofundamental/loops/breakpkg"
	"Ashish/Gofundamental/loops/continuepkg"
	"Ashish/Gofundamental/loops/packagedemo/rectangle"
	"Ashish/Gofundamental/switch/days"
	"TestPackage/Division"
	"TestPackage/Modus"
	"TestPackage/Multiplication"
	"TestPackage/Substarction"
	"fmt"
)

func main() {

	var a float64
	var b float64

	var price, qty int
	var result int
	var result1 float64

	var length, width float64

	var no int

	var strresult string

	var rectLen, rectWidth float64

	var ch int

	choice := 0

	fmt.Println("\nWhich Operation Do you want to perform")

	fmt.Println("1.Addition", "\n2.Substarction", "\n3.Multiplication", "\n4.Division", "\n5.Modus",
		"\n6.Calcualte Total Price using fun", "\n7.Find Area and Perimeter using multivalue fun",
		"\n8.Break example whatever you no enter it only prints 5 no which is condtion we applay",
		"\n9.CHeck No is Even Or odd", "\n10.Find Area of Rectangle and Diagonal of rectangle", "\n11.Know the day of week",
		"\n12.Check No is less graeter or equal to 50")

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

	case 6:
		fmt.Println("Enter Price and Quantity\n")
		fmt.Scanf("%d %d", &price, &qty)

		result = TestFun.CalBill(price, qty)

		fmt.Println("Total price", result)

	case 7:
		fmt.Println("Enter  Length and Width\n")
		fmt.Scanf("%f %f", &length, &width)
		result1, result := multivalue.RectProps(length, width)
		fmt.Println("Area %f  Perimeter  %f", result1, result)

	case 8:
		fmt.Print("Enter No\n")
		fmt.Scanf("%d ", &no)
		result = breakpkg.Funbreak(no)
		fmt.Println("result is ", result)

	case 9:
		fmt.Print("Enter No to be Check even or odd\n")
		fmt.Scanf("%d", &no)
		strresult = continuepkg.Con(no)
		fmt.Print(" ", strresult)

	case 10:
		fmt.Print("Enter length and Width\n")
		fmt.Scanf("%f %f", &rectLen, &rectWidth)
		result1 = rectangle.Area(rectLen, rectWidth)
		fmt.Print("Area of rectangle is\t", result1)
		fmt.Printf("\ndiagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
		fmt.Scanf("%f %f", &rectLen, &rectWidth)

	case 11:
		fmt.Print("Enter 1 t0 7 to know the day of week\n")
		fmt.Scanf("%d", &ch)
		strresult = days.Days(ch)
		fmt.Print("Day is  ", strresult)

	case 12:
		fmt.Print("Enter the Number")
		fmt.Scanf("%d", &no)
		strresult = elseif.Elseifex(no)
		fmt.Print(" ", strresult)

	}

}
