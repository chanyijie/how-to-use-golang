package main

import (
	"fmt"
)

func printInputMultipleTimes(input string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("%s: %d\n", input, i)
	}
}

func main() {
	printInputMultipleTimes("david")
	go printInputMultipleTimes("david-goroutine")

	go func(msg string) {
		fmt.Println("Greeting from an anonymous functioni:", msg)
	}("david")

	var input string
	fmt.Scanln(&input)
}
