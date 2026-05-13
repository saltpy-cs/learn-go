package main

import "fmt"

type contactInfo struct {
	email    string
	postCode string
}

type person struct {
	firstName string
	lastName  string
	age       int
	height    int
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
