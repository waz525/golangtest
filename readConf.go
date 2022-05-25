package main

import (
    "fmt"
    "./Utils"
)

func main() {
    fmt.Println("config read test!")

    ini_parser := Utils.IniParser{}

    conf_file_name := "readConf.ini"
    if err := ini_parser.Load(conf_file_name); err != nil {
        fmt.Printf("try load config file[%s] error[%s]\n", conf_file_name, err.Error())
        return
    }

    Host := ini_parser.GetString("system", "Host")
    Port := ini_parser.GetInt64("system", "Port")
    HostId := ini_parser.GetString("system", "HostId")

    fmt.Printf("Host: %s, Port: %d, HostId:%s\n", Host, Port, HostId)
}
