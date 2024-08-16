package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	fmt.Println("Key from .env = ", os.Getenv("PLATFORM"))
}
