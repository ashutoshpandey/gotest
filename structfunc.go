package main

import "fmt"

type Square struct {
	side int
}

func (square Square) area() int {
	return square.side * square.side
}

func main() {
	square := Square{
		side: 5,
	}

	fmt.Println(square.area())
}
