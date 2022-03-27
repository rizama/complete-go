package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//	Declaring way 1
	sam := person{"Sam", "Pratama"}
	sam.lastName = "Rizama"
	fmt.Println(sam)

	//	Declaring way 2 -> Recomended
	sem := person{firstName: "Sem", lastName: "Pratame"}
	fmt.Println(sem)

	//Updating Struct values
	var rizky person
	rizky.firstName = "Rizky"
	rizky.lastName = "Sam"
	fmt.Println(rizky)
	fmt.Printf("%+v", rizky)
}
