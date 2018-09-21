package main

import (
	"fmt"
)

type person struct {
	age       int16
	firstName string
	lastName  string
	contact   contact
}

type contact struct {
	email   string
	phone   string
	address string
}

func main() {
	var harry person
	harry.age = 32
	harry.firstName = "harry"
	harry.lastName = "lamb"
	harry.contact.email = "harry.yang84@gmail.com"
	cecilia := person{
		firstName: "jiangmei",
		lastName:  "zhao",
		contact: contact{
			email: "riverose@gmail.com",
			phone: "77812315",
		},
	}
	fmt.Printf("%s", cecilia.contact.phone)
}
