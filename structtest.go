package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main4() {
	person := Person{
		Name: "Ashutosh",
		Age:  45,
	}

	show(person)
	fmt.Printf("\n--------------------\n")
	person.ShowDetails()
}

func show(person Person) {
	fmt.Printf("%s , %d", person.Name, person.Age)
}

// Add a function to Person struct
func (person *Person) ShowDetails() {
	fmt.Printf("%s , %d", person.Name, person.Age)
}
