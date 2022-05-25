package main

import (
	"fmt"
)

func isSuShu(n int) bool {
	rst := true
	s1 := n/2 + 1
	for i := 2; i < s1; i++ {
		if n%i == 0 {
			rst = false
			break
		}
		s1 = n/i + 1
	}
	return rst
}

func getSushuNum(max int) int {
	count := 0
	for i := 2; i <= max; i++ {
		if isSuShu(i) {
			//fmt.Println(i)
			count++
		}
	}
	return count
}

func main() {

	rst := getSushuNum(100000000)
	//rst := getSushuNum(10000)
	fmt.Println(rst)

}

