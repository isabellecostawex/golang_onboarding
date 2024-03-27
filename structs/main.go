package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	jim := person{
	firstName: "Jim", 
	lastName: "Party",
	contact: contactInfo{
		email: "jim@gmail.com",
		zipCode: 954400,
	},
}
	//JimPointer := &jim
	//JimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
