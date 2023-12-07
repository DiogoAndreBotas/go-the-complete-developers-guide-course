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
	diogo := person{
		firstName: "Diogo",
		lastName:  "Botas",
		contactInfo: contactInfo{
			email:   "diogo.andre.botas@gmail.com",
			zipCode: 18000,
		},
	}

	diogo.updateName("Josefino")
	diogo.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}
