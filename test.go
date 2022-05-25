package main

import (
		"fmt"
//		"strings"
//		"net/url"
		"encoding/json"
		"./Utils"
)
type ConnCount struct {
		Hostid	string	`json:"Hostid"`
		Ruleid	string	`json:"Ruleid"`
		Conn	string	`json:"Conn"`
}

func main() {
//		fmt.Println(strings.EqualFold("1HELLO", "hello"))
//		rst := Utils.HttpGet("http://121.40.96.81:12800/filestatic/TargeFile")
//		fmt.Println(rst)

		//data := url.Values{"app_id":{"238b2213-a8ca-42d8-8eab-1f1db3c50ed6"}, "mobile_tel":{"13794227450"}}

//		data := "[{\"Hostid\":\"10\",\"Ruleid\":\"10\",\"Conn\":\"10\"},{\"Hostid\":\"11\",\"Ruleid\":\"11\",\"Conn\":\"11\"}]"
		data := `[{"Hostid":"10","Ruleid":"10","Conn":"10"},{"Hostid":"11","Ruleid":"11","Conn":"11"}]`
		rst := Utils.HttpPost("http://121.40.96.81:12800/ConntrackServer/connCount",data,"application/x-www-form-urlencoded")
		fmt.Println(rst)

		//data = `{"Hostid":"10","Ruleid":"10","Conn":"10"}`
		b := []byte(data)
		nConn := []ConnCount{}


		if err := json.Unmarshal(b, &nConn); err == nil {
				fmt.Println(nConn)
		} else {
				fmt.Println(err)
		}
}
