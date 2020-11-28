package main

import "fmt"

func main() {
	result := add(2, 3)
	fmt.Println(result)

	mult := multiply(4, 5)
	fmt.Println(mult)

	div, mod := divmod(7, 2)
	fmt.Printf("div=%d, mod = %d \n", div, mod)

}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}
