package main

import "fmt"

func main() {
	// STRING
	var myString string
	fmt.Println(myString) // prints nothing
	myString = "Welcome to coding."
	fmt.Println(myString) // prints the string

	// INTEGER
	var myInt int
	fmt.Println(myInt) // prints 0
	myInt = 10
	fmt.Println(myInt)

	// BOOLEAN
	// zero value for bool is false
	var myBool bool
	fmt.Println(myBool)
	myBool = true
	fmt.Println(myBool)

	// UNSIGNED INTEGER
	var unsignedInteger uint8
	fmt.Println(unsignedInteger)
	unsignedInteger = 10
	// unsignedInteger = -10 // Error : overflows
	print(unsignedInteger)

	// TYPECAST
	var myUnsignedInteger = -10
	var myInteger int = 230000234
	fmt.Println(myInteger)
	myInteger = int(myUnsignedInteger)
	fmt.Println((myInteger))

	// FLOAT
	var myFloat float32
	fmt.Println(myFloat)
	myFloat = 23.423433523432432 // Output: 23.423433
	fmt.Println(myFloat)

	var myBigFloat float64
	fmt.Println(myBigFloat)
	myBigFloat = 43.22133124232343241343413431413 // Output: 43.22133124232343
	fmt.Println(myBigFloat)

	// COMPLEX
	var myComplex complex128
	myComplex = complex(myBigFloat, myBigFloat)
	fmt.Println(myComplex)

	var myRealPart, myImaginaryPart float64
	myRealPart = real(myComplex)
	myImaginaryPart = imag(myComplex)
	fmt.Println(myRealPart)
	fmt.Println(myImaginaryPart)

	// STRING

}
