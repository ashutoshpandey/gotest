package main

import (
	"fmt"
)

type Dog struct {
	Name string
}

func (d Dog) Bark() {
	fmt.Printf("\n%s barking", d.Name)
	d.Name = "Champa"
}

func (d *Dog) Walk() {
	fmt.Printf("\n%s walking", d.Name)
	d.Name = "Dhamaka"
}

func main6() {
	d1 := Dog{"Tommy"}
	d1.Bark()
	d1.Walk()
	d1.Bark()
}
