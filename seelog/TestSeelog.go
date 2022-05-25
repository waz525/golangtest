// TestSeelog.go
package main

import (
    seelog "github.com/cihub/seelog"
)

func main() {
    logger, err := seelog.LoggerFromConfigAsFile("./logconfig.xml")

    if err != nil {
        seelog.Critical("err parsing config log file", err)
        return
    }
    seelog.ReplaceLogger(logger)

    seelog.Error("seelog error")
    seelog.Info("seelog info")
    seelog.Debug("seelog debug")
	seelog.Flush()
}

