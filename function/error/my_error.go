package main

import (
	"fmt"
	"math"
)

func main() {
	s1, err := sqrt(2.0)
	print(s1, err)

	s2, err := sqrt(-2.0)
	print(s2, err)
}

func print(data float64, err error) {
	if err != nil {
		fmt.Printf("ERROR :%s\n", err)
	} else {
		fmt.Printf("result is %f\n", data)
	}
}

func sqrt(number float64) (float64, error) {
	if number < 0 {
		return 0.0, fmt.Errorf("sqrt of negative value %f ", number)
	}

	return math.Sqrt(number), nil
}
