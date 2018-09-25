package main

import (
	"fmt"
)

func main() {
	nickNames := map[string]string{
		"Sam Cassell":     "ET",
		"Hakeem Olajuwon": "The Dream",
	}
	fmt.Println(nickNames["Olajuwon"])
}
