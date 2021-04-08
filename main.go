package main

import "fmt"

func Calculate(num int) (result int) {
	result = num * 2
	return result
}

func main() {
	fmt.Printf("Calculate\n")
	result := Calculate(2)
	fmt.Printf("result : %d", result)
}
