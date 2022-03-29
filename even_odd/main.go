package main

import "fmt"

type numbersInt []int

func myPrint(n numbersInt) {
	// Looping for Check the number is even or odd
	for _, number := range n {
		// Check the number
		if number%2 == 0 {
			//	Print the result
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}

func newNumbers(max int) numbersInt {
	numbers := numbersInt{}

	//	generate new numbers
	for i := 0; i <= max; i++ {
		numbers = append(numbers, i)
	}

	return numbers
}

func main() {
	//	Create Slice
	numbers := newNumbers(10)

	//	Print even and odd
	numbers.myPrint2()
}

func (n numbersInt) myPrint2() {
	// Looping for Check the number is even or odd
	for _, number := range n {
		// Check the number
		if number%2 == 0 {
			//	Print the result
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
