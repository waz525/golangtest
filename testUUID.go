package main

import (
    "fmt"
    "log"
    "os/exec"
)

func main() {
    out, err := exec.Command("uuidgen | sed 's/-//g' | cut -b 1-24").Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s", out)
}
