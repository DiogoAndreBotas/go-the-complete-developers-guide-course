package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	diogo := person{
		firstName: "Diogo",
		lastName:  "Botas",
		contact: contactInfo{
			email:   "diogo.andre.botas@gmail.com",
			zipCode: 18000,
		},
	}

	fmt.Printf("%+v", diogo)
}
