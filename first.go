package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main1() {
	a := "Hello world from variable"
	b, c := 5, "Good"
	var d int
	var e int = 6
	const f = 10
	var (
		g = 1
		h = false
		i = "good"
	)

	d = 5

	fmt.Println(a, b)
	fmt.Printf("C = %s\n", c)
	fmt.Println(d, e, f)
	fmt.Println(g, h, i)

	fmt.Println(time.Now())

	rnd := rand.Intn(100)
	fmt.Println(rnd)

	fmt.Println(add(4, 5))

	var result int = multiply(4, 5)
	fmt.Println(result)
}

func add(a, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}
