package main

import "fmt"

func test() (int,int) {
		return 2,3
}

func main() {
    fmt.Println("Hello, World!")
	i,j := test()
	fmt.Println(i,j)
}
