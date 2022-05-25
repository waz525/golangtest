package main

import (
	"fmt"
)

func main() {
	var inList []int
	inList = append(inList , 1 )
	inList = append(inList , 2 )
	inList = append(inList , 3 )

	outList := inList[0:0]
	fmt.Println(len(outList))

}

