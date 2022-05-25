package main

import (
	"fmt"
	"./Utils"
	"time"
)

func RunTest(i int,pool *Utils.Pool)  {
	fmt.Println(i)
	time.Sleep(100 * time.Millisecond)
	pool.Done()
}

func main()  {
	size:=2
	pool := Utils.NewPool(size)
	for i:=1;i<=10;i++{
		pool.Add(1)
		go RunTest(i,pool)
	}
	pool.Wait()
}
