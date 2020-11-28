package main

import "fmt"

func main() {
	stocks := map[string]float64{
		"AMZN": 1232.3,
		"GOOG": 1112.4,
		"MSFT": 23.45,
	}

	fmt.Println(len(stocks))

	//get value
	fmt.Printf("value of MSFT is %f \n", stocks["MSFT"])

	//get not found data
	fmt.Printf("value of RAND is %f \n", stocks["RAND"])

	//print if found, otherwise print not found error
	key := "AMZN"
	value, hasdata := stocks[key]
	if !hasdata {
		fmt.Printf("data not found \n")
	} else {
		fmt.Printf("data found for key %s -> %v \n", key, value)
	}

	//insert map
	stocks["TSLA"] = 2323.34
	fmt.Printf("value of TSLA is %f \n", stocks["TSLA"])

	//delete
	delete(stocks, "TSLA")
	fmt.Printf("value of TSLA after delete is %f \n", stocks["TSLA"])

	//iterate
	for key, value := range stocks {
		fmt.Printf("value of %s is %f \n", key, value)
	}

}
