package main

import (
	"flag"
	"fmt"
	"strings"
	"net/http"
	"net/http/cgi"
)

var CigBinPrefix string
var CigBinPath   string
var FileStaticPath string
var ServerPort		string
var FileStaticPrefix string

//默认页面处理，用于支持cgi
func DealWithRequestForCgi(rw http.ResponseWriter, req *http.Request) {
		url_path := req.URL.Path
		fmt.Println("Receive "+req.Method+" from "+req.RemoteAddr+" ,URL.Path: "+url_path)
		//判断是否是cgi-bin请求
		if strings.Index(url_path, CigBinPrefix) == 0 {
				handler := new(cgi.Handler)
				handler.Path = CigBinPath+"/"+url_path[len(CigBinPrefix)-1:]
				fmt.Println("handler.Path: "+handler.Path)
				handler.Dir = CigBinPath
				handler.ServeHTTP(rw, req)
		} else {
				http.NotFound(rw, req)
		}
}

//开启Http服务
func StartHttpService() {
		http.HandleFunc("/", DealWithRequestForCgi)

		//设置静态文件路径
		fs := http.FileServer(http.Dir(FileStaticPath))
		http.Handle(FileStaticPrefix, http.StripPrefix(FileStaticPrefix, fs))

		fmt.Println("Server Listen on "+ServerPort+" ...")
		//开启http侦听
		err := http.ListenAndServe(":"+ServerPort, nil)
		if err != nil {
				fmt.Println("ListenAndServe Error:", err)
		}
}


func main() {
		flag.StringVar(&ServerPort, "p", "8000", "The Port for Httpd .")
		flag.StringVar(&FileStaticPath, "w", "/var/www/html/" , "The Path of WebSit .")
		flag.StringVar(&FileStaticPrefix, "s", "/myweb/", "The URL of WebSit .")
		flag.StringVar(&CigBinPath, "c", "/var/www/cgi-bin/", "The Path of Cgi-bin ." )
		flag.StringVar(&CigBinPrefix, "b", "/myweb/cgi-bin/", "The URL Path for Cgi-bin .")
		flag.Parse()

		StartHttpService()
}
