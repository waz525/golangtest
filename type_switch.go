package main

import "fmt"

func main() {
	 var x interface{}

	 switch i := x.(type) {
		case nil:
		 fmt.Println("x 的类型 :%T",i)
		case int:
		 fmt.Println("x 是 int 型")
		case float64:
		 fmt.Println("x 是 float64 型")
		case func(int) float64:
		 fmt.Println("x 是 func(int) 型")
		case bool, string:
		 fmt.Println("x 是 bool 或 string 型" )
		default:
		 fmt.Println("未知型")
	 }
}

