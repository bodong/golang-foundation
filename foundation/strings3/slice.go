package main

import "fmt"

func main() {
	//same type
	datas := []string{"one", "two", "three"}
	fmt.Printf("data : %v | type of %T \n", datas, datas)

	//length
	fmt.Println(len(datas))

	//print using specific index
	fmt.Printf("data index 0 is %v \n", datas[0])
	fmt.Printf("data slices are %v \n", datas[1:])

	//print using loop by index
	for i := 0; i < len(datas); i++ {
		fmt.Printf(" index %d -> %v \n", i, datas[i])
	}

	//print using loop range
	for i, data := range datas {
		fmt.Printf(" range -> index %d -> %v \n", i, data)
	}

	//print using loop range, but ignore the index
	for _, data := range datas {
		fmt.Printf(" range -> %v \n", data)
	}
}
