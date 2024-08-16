package main

import (
	"fmt"
	mods "gotest/modules"
)

func main3() {
	a := 5
	b := 7
	result := mods.Jodo(a, b)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, result)
}
