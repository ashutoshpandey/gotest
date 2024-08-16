package main

import (
	"fmt"
)

func main2() {
	// Regular for loop
	var n int = 5
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%d x %d = %d\n", n, i, n*i)
		}
	}

	// Loop on array
	//ar := [4]int{10, 20, 30, 40}
	//ar := [...]int{10, 20, 30, 40}
	ar := []int{10, 20, 30, 40}
	for _, value := range ar {
		fmt.Println(value)
	}

	// Loop on map
	users := map[string]int{
		"Ashu":  10,
		"Shiva": 20,
	}

	for key, value := range users {
		fmt.Printf("%s %d\n", key, value)
	}

	// Another way
	data := make(map[string]int)
	data["Ashu"] = 30
	data["Shiva"] = 40

	for key, value := range users {
		fmt.Printf("%s %d\n", key, value)
	}
}
