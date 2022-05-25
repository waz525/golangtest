// convert json to struct

package main


import (
	"fmt"
	"encoding/json"
)

/*
type Host struct {
	IP string
	Name string
}
*/
type ConnCount struct {
        Hostid  string  `json:"hostid"`
        Ruleid  string  `json:"ruleid"`
        Conn    string  `json:"conn"`
}


func main() {

	//b := []byte(`{"IP": "192.168.11.22", "name": "SKY"}`)
	b := []byte(`{"hostid":"10","ruleid":"10","conn":"10"}`)

	m := ConnCount{}

	err := json.Unmarshal(b, &m)
	if err != nil {

		fmt.Println("Umarshal failed:", err)
		return
	}


	fmt.Println("m:", m)
//	fmt.Println("m.IP:", m.IP)
//	fmt.Println("m.Name:", m.Name)
}
