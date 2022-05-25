package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    var s =  map[string]interface{}{}
    var a  = map[string]interface{}{"b":11111}

    s["nihao"] = map[string]interface{}{"nihao":"dddd","bb":a}
    res,_ := json.Marshal(s)
    resString := string(res)
    fmt.Println(resString)
}

