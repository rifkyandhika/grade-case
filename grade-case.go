package main

import (
	"fmt"
)

func main() {
	fmt.Print("Input Grade: ")
	var input int
	fmt.Scan(&input)

	if input >= 0 && input <= 100 {
		if input >= 38 {
			multiple := (input + 4) / 5 * 5

			difference := multiple - input

			if difference < 3 {
				fmt.Printf("Final Grade: %d\n", multiple)
			} else {
				fmt.Printf("Final Grade: %d\n", input)
			}
		} else {
			fmt.Printf("Final Grade: %d\n", input)
		}
	} else {
		fmt.Println("Error: Invalid input. Please enter a valid non-negative integer or not exceeding 100.")
	}
}
