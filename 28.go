package main
import  (
    "fmt"
    "sort"
)

func main() {
    var nCount, nFlag int
    fmt.Scan(&nCount)
    nList := make([]int, nCount)
    for i:=0;i<nCount;i++ {
        fmt.Scan(&nList[i])
    }
    fmt.Scan(&nFlag)
    if nFlag == 0 {
        //升序
        sort.Ints(nList)
    } else {
        //降序
        sort.Sort(sort.Reverse(sort.IntSlice(nList)))
    }
    for i:=0;i<nCount;i++ {
        fmt.Print(nList[i]," ")
    }
	fmt.Println()
}
