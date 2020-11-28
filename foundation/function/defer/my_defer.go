package main

import "fmt"

func main() {
	worker()
}

func worker() {
	defer cleanup("A")
	defer cleanup("B")
	defer cleanup("D")
	defer cleanup("E")
	defer cleanup("F")
	fmt.Println("worker do the work")
}

func cleanup(resource string) {
	fmt.Printf("cleanup resource %s \n", resource)
}
