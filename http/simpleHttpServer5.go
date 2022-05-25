/*
CGI包使用
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cgi"
)


func test(rw http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL.Path)
		handler := new(cgi.Handler)
		/*
		/var/www/cgi-bin/test.sh :

		#!/bin/sh
		printf "Content-Type: text/plain;charset=utf-8\n\n"
		env

		*/
		handler.Path = "/var/www/cgi-bin/test.sh"
		fmt.Println("handler.Path: ",handler.Path)
		handler.Dir = "/var/www/cgi-bin/"
		handler.ServeHTTP(rw, req)
}

func indexpage(rw http.ResponseWriter, req *http.Request) {
		fmt.Println(req.URL.Path)
	fmt.Fprintln(rw,"This is Index Page")
}

func main() {
	http.HandleFunc("/cgi-bin/test", test)
	http.HandleFunc("/",indexpage)
	log.Fatal(http.ListenAndServe(":12802", nil))
}
