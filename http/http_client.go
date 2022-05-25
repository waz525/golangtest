package main

import (
    "bytes"
//    "encoding/json"
    "io"
    "io/ioutil"
    "net/http"
    "net/url"
    "time"
    "fmt"
    "strings"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

    // 超时时间：5秒
    client := &http.Client{Timeout: 5 * time.Second}
    resp, err := client.Get(url)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    var buffer [512]byte
    result := bytes.NewBuffer(nil)
    for {
        n, err := resp.Body.Read(buffer[0:])
        result.Write(buffer[0:n])
        if err != nil && err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }
    }

    return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data url.Values, contentType string) string {

    // 超时时间：5秒
    client := &http.Client{Timeout: 5 * time.Second}
    //jsonStr, _ := json.Marshal(data)
    resp, err := client.Post(url, contentType, strings.NewReader(data.Encode()))
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    result, _ := ioutil.ReadAll(resp.Body)
    return string(result)
}

func main() {
	fmt.Println(Get("http://www.01happy.com/demo/accept.php?id=1"))
	fmt.Println("-----------------------------")
//	data := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	data := url.Values{"app_id":{"238b2213-a8ca-42d8-8eab-1f1db3c50ed6"}, "mobile_tel":{"13794227450"}}
	fmt.Println(Post("http://www.01happy.com/demo/accept.php",data,"application/x-www-form-urlencoded"))
}
