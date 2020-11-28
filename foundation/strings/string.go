package main

import "fmt"

func main() {
	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])

	//String is immutable
	//book[0] = 12

	//slice
	slice1 := book[3:8]
	fmt.Println(slice1)

	slice2 := book[:7]
	fmt.Println(slice2)

	slice3 := book[2:]
	fmt.Println(slice3)

	//concat
	concat := "the book of " + slice2
	fmt.Println(concat)

	//multiline
	longstring := `
	today 
	is the best day 
	to plan everything
	from zero again
	....
	`
	fmt.Println(longstring)
}
