package main

import "fmt"

type numbersInt = []int

func main() {
	//	Create Slice 0-10
	numbers := newNumbers(10)

	// Looping for Check the number is even or odd
	for _, number := range numbers {
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
