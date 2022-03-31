package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	//	Declaring way 1 without embedded another struct
	//sam := person{"Sam", "Pratama"}
	//sam.lastName = "Rizama"
	//fmt.Println(sam)

	//	Declaring way 2 without embedded another struct -> Recomended
	//sem := person{firstName: "Sem", lastName: "Pratame"}
	//fmt.Println(sem)

	//Updating Struct values without embedded another struct
	//var rizky person
	//rizky.firstName = "Rizky"
	//rizky.lastName = "Sam"
	//fmt.Println(rizky)
	//fmt.Printf("%+v", rizky)

	sampratama := person{
		firstName: "SAM",
		lastName:  "PRATAMA",
		contactInfo: contactInfo{
			email:   "sam@gmail.com",
			zipCode: 12345,
		},
	}

	sampratama.myPrint()
	samPointer := &sampratama
	samPointer.updateName("Rizama")
	sampratama.myPrint()
}

func (p *person) updateName(newFirstname string) {
	(*p).firstName = newFirstname
}

func (p person) myPrint() {
	fmt.Printf("%+v", p)
}
