package main

import (
	"clifolio/internal/ui"
	"fmt"

	"github.com/joho/godotenv"
)



func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Oh no! env file not found.")
	}
	ui.App()
}
