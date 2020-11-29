package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	/* 	<-ch
	   	fmt.Println("here")
	*/
	go func() {
		ch <- 34
	}()

	// receive channel
	val := <-ch
	fmt.Printf("%d\n", val)
	fmt.Println("---------")

	//send multiple
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	fmt.Println("---------")

	//close to signal
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()
	for i := range ch {
		fmt.Printf("received %d\n", i)
	}

}
