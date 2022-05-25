package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"os"
	"strconv"
	"path"
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
		log.Println("Receive "+req.Method+" from "+req.RemoteAddr+" ,URL.Path: "+url_path)
		//判断是否是cgi-bin请求
		if strings.Index(url_path, CigBinPrefix) == 0 {
				handler := new(cgi.Handler)
				handler.Path = CigBinPath+"/"+url_path[len(CigBinPrefix)-1:]
				log.Println("handler.Path: "+handler.Path)
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

		log.Println("Server Listen on "+ServerPort+" ...")
		//开启http侦听
		err := http.ListenAndServe(":"+ServerPort, nil)
		if err != nil {
				log.Println("ListenAndServe Error:", err)
		}
}

func init() {
		pid := os.Getpid()
	    log.SetPrefix("[PID:"+strconv.Itoa(pid)+"] ")
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		//log.SetFlags(log.Ldate | log.Ltime )
}


func main() {
		flag.StringVar(&ServerPort, "p", "18000", "The Port for Httpd .")
		flag.StringVar(&FileStaticPath, "w", "/var/www/html/" , "The Path of WebSit .")
		flag.StringVar(&FileStaticPrefix, "s", "/myweb/", "The URL of WebSit .")
		flag.StringVar(&CigBinPath, "c", "/var/www/cgi-bin/", "The Path of Cgi-bin ." )
		flag.StringVar(&CigBinPrefix, "b", "/myweb/cgi-bin/", "The URL Path for Cgi-bin .")
		flag.Parse()

		if os.Getppid() != 1 {

				createLogFile := func (fileName string) (fd *os.File, err error) {
						dir := path.Dir(fileName)
						if _, err = os.Stat(dir); err != nil && os.IsNotExist(err) {
								if err = os.MkdirAll(dir, 0755); err != nil {
										fmt.Printf("Start-Daemon: create dir: %s failed, err is: %v\n", dir, err)
										return
								}
						}
						if fd, err = os.OpenFile(fileName,os.O_RDWR|os.O_CREATE|os.O_APPEND,os.ModeAppend) ; err != nil {
						//if fd, err = os.Create(fileName); err != nil {
								fmt.Printf("Start-Daemon: open log file: %s failed, err is: %v\n", fileName, err)
								return
						}
						return
				}

				logFd, err := createLogFile("/var/log/goHttpd.log")
				if err != nil {
						return
				}
				defer logFd.Close()

				cmdName := os.Args[0]
				newProc, err := os.StartProcess(cmdName, os.Args, &os.ProcAttr{Files: []*os.File{logFd, logFd, logFd}})
				if err != nil {
						log.Println("Start-Deamon: start process: ", cmdName," failed, err is: ", err)
						return
				}
				log.Println("Start-Deamon: run in daemon success, pid:", newProc.Pid)
				return
		}

		log.Println("Config ServerPort:", ServerPort)
		log.Println("Config FileStaticPrefix:", FileStaticPrefix)
		log.Println("Config FileStaticPath:", FileStaticPath )
		log.Println("Config CigBinPrefix:", CigBinPrefix )
		log.Println("Config CigBinPath:", CigBinPath)
		//log.Println("Config :",)




		StartHttpService()
}
