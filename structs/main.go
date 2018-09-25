package main

import (
	"fmt"
)

type person struct {
	firstName string
	LastName  string
	contact   contact
}

type contact struct {
	emailAddr string
}

func main() {
	jason := person{
		firstName: "Jason",
		LastName:  "Smith",
		contact: contact{
			emailAddr: "jason.smith@abc.com",
		},
	}
	fmt.Printf("%s", jason.contact.emailAddr)
	jason.print()
	jason.updateFirstName("Harry")
}

func (p person) print() {
	fmt.Printf("%s %s", p.firstName, p.LastName)
}

func (p *person) updateFirstName(newFirst string) {
	(*p).firstName = newFirst
	(*p).print()
}
