package main

import "fmt"

func main() {
   /* 数组 - 5 行 2 列*/
   var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
   var i, j int

   /* 输出数组元素 */
   for  i = 0; i < 5; i++ {
      for j = 0; j < 2; j++ {
         fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
      }
   }


    // 创建二维数组
    sites := [2][2]string{}

    // 向二维数组添加元素
    sites[0][0] = "Google"
    sites[0][1] = "Runoob"
    sites[1][0] = "Taobao"
    sites[1][1] = "Weibo"

    // 显示结果
    fmt.Println(sites)



    // Step 1: 创建数组
    values := [][]int{}

    // Step 2: 使用 appped() 函数向空的二维数组添加两行一维数组
    row1 := []int{1, 2, 3}
    row2 := []int{4, 5, 6}
    values = append(values, row1)
    values = append(values, row2)

    // Step 3: 显示两行数据
    fmt.Println("Row 1")
    fmt.Println(values[0])
    fmt.Println("Row 2")
    fmt.Println(values[1])

    // Step 4: 访问第一个元素
    fmt.Println("第一个元素为：")
    fmt.Println(values[0][0])
}
