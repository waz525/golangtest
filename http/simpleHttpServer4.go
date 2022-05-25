/*
多端口侦听
*/
package main

import (
	"fmt"
	"net/http"
)

func index1(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("This is server 1 ."))
}

func index2(w http.ResponseWriter, r *http.Request) {
    _, _ = w.Write([]byte("This is server 2 ."))
}


func http_server1() {
		server1 := http.NewServeMux()
		server1.HandleFunc("/", index1)
		err := http.ListenAndServe(":12801", server1)
		if err != nil {
			fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
			return
		}
		fmt.Println("Listen 12801 ...")
}

func http_server2() {
        server2 := http.NewServeMux()
        server2.HandleFunc("/", index2)
		err := http.ListenAndServe(":12802", server2)
        if err != nil {
            fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
            return
        }
		fmt.Println("Listen 12802 ...")

}

func main() {
		//以多线程方式开启server1
		go http_server1()
		//主进程开户server2, 主进程必须常驻运行，否则程序退出
		http_server2()
}

