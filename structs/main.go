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
	// samPointer := &sampratama // Give me the memory address of the values this variable is pointing at
	sampratama.updateName("Rizama")
	sampratama.myPrint()
}

//*person :this is a description type - it means we're working with a pointer to a person
//*p :this is a operator - it means we want to manipulate the values pointer is referencing
func (pointerToPerson *person) updateName(newFirstname string) {
	(*pointerToPerson).firstName = newFirstname //Give me the value this memory address is pointing at
}

func (p person) myPrint() {
	fmt.Printf("%+v", p)
	fmt.Println()
}
