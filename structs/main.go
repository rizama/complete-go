package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//	Declaring way 1
	sam := person{"Sam", "Pratama"}
	fmt.Println(sam)

	//	Declaring way 2 -> Recomended
	sem := person{firstName: "Sem", lastName: "Pratame"}
	fmt.Println(sem)
}
