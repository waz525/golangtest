/*
解析Content-Type为 application/json的post请求体参数
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
//	"io/ioutil"
)

type person_struct struct {
		FirstName string `json:"FirstName"`
		LastName  string `json:"LastName"`
}

/*
//读取json传递数据  data:'{"FirstName":"yd","LastName":"123456"}',
func test(rw http.ResponseWriter, req *http.Request) {
        fmt.Println("method:", req.Method)

        if req.Method != "POST" && req.Method != "post" {
                fmt.Fprintln(rw, "aaaa")
                return
        }

    var t person_struct

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        fmt.Printf("read body err, %v\n", err)
        return
    }
    fmt.Println("req.Body: ", string(body) )

    //decoder := json.NewDecoder(req.Body)
    //err = decoder.Decode(&t)
    err = json.Unmarshal(body,&t)
    if err != nil {
        panic(err)
    }

    jsonBytes, err := json.Marshal(t)
    if err != nil {
       fmt.Println(err)
    }
    fmt.Println(string(jsonBytes))

    rw.Header().Set("Access-Control-Allow-Origin", "*")
    rw.Header().Set("Content-Type", "application/json;charset=utf-8")
    fmt.Fprintln(rw, string(jsonBytes) )



*/


func test(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("method:", req.Method)

		if req.Method != "POST" && req.Method != "post" {
				fmt.Fprintln(rw, "aaaa")
				return 
		} 

	var t person_struct

	t.FirstName = req.FormValue("FirstName")
	t.LastName = req.FormValue("LastName")
	//转成json
	jsonBytes, err := json.Marshal(t)
    if err != nil {
	   fmt.Println(err)
    }
    fmt.Println(string(jsonBytes))

	rw.Header().Set("Access-Control-Allow-Origin", "*")
    rw.Header().Set("Content-Type", "application/json;charset=utf-8")
	fmt.Fprintln(rw, string(jsonBytes) )
}

func indexpage(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw,"This is Index Page")
}

func main() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/",indexpage)
	log.Fatal(http.ListenAndServe(":12802", nil))
}
