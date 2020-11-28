package main

import "fmt"

func main() {

	//count = 0
	count := 0

	//for every pair of 4 digit numbers
	for a := 1; a <= 9; a++ {
		for b := a; b <= 9; b++ { //so will not count twice
			n := a * b

			// if a * b even ended
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				// count = count + 1
				count++
			}
		}
	}

	fmt.Println(count)

}
