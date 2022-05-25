package crypt

import (
    "fmt"
    "testing"
)

func Test0001(test *testing.T) {
    s := Encrypt("1234", "1234")
    fmt.Println("TestEncrypt = ", s)
}

func Test0002(test *testing.T) {
    s := Decrypt("1234", "1234")
    fmt.Println("TestDecrypt = ", s)
}
