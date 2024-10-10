package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func main() {
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       25,
	}

	fmt.Println("Full Name:", person.FullName())

	if person.IsAdult() {
		fmt.Println(person.FullName(), "is an adult.")
	} else {
		fmt.Println(person.FullName(), "is not an adult.")
	}
}
