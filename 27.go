package main

import (
	"fmt"
)

func running(i int, ch chan int) {
	fmt.Println(i)
	ch <- i
}
func main() {

	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go running(i, ch)
	}
    for i := 0; i < 10; i++ {
		<-ch
	}
	fmt.Println("Main End")
}

