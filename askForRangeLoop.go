package main

import "fmt"

func main() {
	// result := 1
	// for i := 1; i < 5; i++ {
	// 	result *= i
	// }

	// value as index
	// call for range

	names := []string{"Wen", "Shu", "Bruce"}
	for _, s := range names {
		fmt.Println(s)
	}

	// fmt.Println("Results", result)
}
