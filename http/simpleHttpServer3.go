/*
参数读取
*/
package main

import (
    "fmt"
    "net/http"
)

func main() {
	// http://x.x.x.x:12800/?name=234&email=123@123.com
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome to my website!")
		fmt.Fprintln(w, r.URL.Query().Get("name"))
		fmt.Fprintln(w, r.FormValue("email"))
    })

	// http://x.x.x.x:12800/static/
    fs := http.FileServer(http.Dir("."))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.ListenAndServe(":12800", nil)
}
