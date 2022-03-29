package main

import "fmt"

func main() {
	result := 1
	for i := 1; i < 5; i++ {
		result *= i
	}

	fmt.Println("Results", result)
}
