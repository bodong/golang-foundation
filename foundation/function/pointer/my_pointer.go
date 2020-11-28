package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4}
	fmt.Printf("Original data %v \n", values)
	//by reference
	doubleAt(values, 3)
	fmt.Println(values)

	//by value
	num := 4
	double(num)
	fmt.Println(num)

	//by pointer / reference
	doublePtr(&num)
	fmt.Println(num)

}

func doublePtr(num *int) {
	*num *= 2
}

func double(num int) {
	num *= 2
}

func doubleAt(values []int, i int) {
	values[i] *= 2
}
