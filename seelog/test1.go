package main

import (
    seelog "github.com/cihub/seelog"
)

func main() {
		seelog.Error("seelog error")
		seelog.Info("seelog info")
		seelog.Debug("seelog debug")
		seelog.Flush()
}
