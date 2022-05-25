//命令行参数： XXXXX -u root -p 123456 -P 3370 -h 127.0.0.1
package main
 
import (
    "flag"
    "fmt"
)
 
func main() {
    // 定义几个变量，用于接收命令行的参数值
    var user        string
    var password    string
    var host        string
    var port        int
    // &user 就是接收命令行中输入 -u 后面的参数值，其他同理
    flag.StringVar(&user, "u", "root", "账号，默认为root")
    flag.StringVar(&password, "p", "", "密码，默认为空")
    flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
    flag.IntVar(&port, "P", 3306, "端口号，默认为3306")
	
    // 解析命令行参数写入注册的flag里
    flag.Parse()
    // 输出结果
    fmt.Printf("user：%v\npassword：%v\nhost：%v\nport：%v\n",
        user, password, host, port)
}
