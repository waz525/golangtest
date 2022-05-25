/* 读取文件 */
package main


import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    ReadLine2("19.txt")
}

func ReadLine2(filename string) {
    f, _ := os.Open(filename)
    defer f.Close()
    r := bufio.NewReader(f)
    for {
        aa, err := readLine(r)
        if err != nil {
            break
        }
        fmt.Println(string(aa))
    }

}

func readLine(r *bufio.Reader) (string, error) {
    line, isprefix, err := r.ReadLine()
    for isprefix && err == nil {
        var bs []byte
        bs, isprefix, err = r.ReadLine()
        line = append(line, bs...)
    }
    return string(line), err
}
