package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

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

	var panton person
	panton.firstName = "panton"
	panton.lastName = "pmanton"

	anton := person{
		firstName: "anton",
		lastName:  "panton",
		contactInfo: contactInfo{
			email:   "anton@anton.com",
			zipCode: 12344,
		},
	}

	// vanton := person{}

	fmt.Println(panton)
	fmt.Println(panton.firstName)
	fmt.Println(panton.lastName)
	fmt.Println(anton)
	fmt.Printf("%+v\n", anton)
	anton.print()

}

func (p person) print() {
	spew.Dump(p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}