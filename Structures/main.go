package main

import "fmt"

type contactinfo struct {
	email string
	zip   int
}
type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

func main() {
	/*METHOD 3
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"*/
	/*METHOD 1
	alex := peron{firstName:"Alex",lastName:"Anderson"}
	In above method you can assign in order and is a good method
	METHOD2
	alex := person{"Alext", "Anderson"}
	//Assigning the values to struct will be based on its order*/

	//fmt.Printf("%+v", alex) //%+v will give the field names for alex
	jim := person{
		firstName: "Jim",
		lastName:  "party",
		contact: contactinfo{
			email: "jim@gmail.com",
			zip:   94000,
		},
	}
	jimPointer := &jim
	jimPointer.updateName("lara")
	jim.print()
}

func (pointerToperson *person) updateName(newFirstName string) {
	(*pointerToperson).firstName = newFirstName
	/*EXPLANATION : *person is not for giving value its just when you call
	the pointerToperson this address will be accessible
	and the (*pointerToperson) is just that we want to update the value
	of it at that address */
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
