package main

import (
	"fmt"
)

func main() {
	/* 	vals := []int{1, 2, 3}
	   	//will make panic
	   	v := vals[6]
	   	fmt.Println(v)
	*/

	/* file, err := os.Open("no-such-file")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fmt.Println(file) */
	v := safeValue([]int{1, 2, 3}, 10)
	fmt.Println(v)

}

func safeValue(vals []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	}()

	return vals[index]
}
