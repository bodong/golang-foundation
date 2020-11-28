package main

import "fmt"

func main() {
	//define by position
	t1 := Trade{"MSFT", 2, 34.36, false}
	fmt.Println(t1)
	fmt.Printf("t1 : %+v \n", t1)

	fmt.Println(t1.Symbol)

	//define by fields
	t2 := Trade{
		Symbol: "MSFT",
		Volume: 5,
		Price:  23.45,
		Buy:    true,
	}
	fmt.Printf("t2 : %+v \n", t2)

	t4 := t2.Value()
	fmt.Printf("t4 : %+v \n", t4)

	//create empty object
	t3 := Trade{}
	fmt.Printf("t3 : %+v \n", t3)

}

//Trade object
type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

//Value calculate actual value
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}
	return value
}
