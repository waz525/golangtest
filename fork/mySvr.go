package main

import (
		"os"
		"fmt"
		"time"

)



func main() {
      // when os.Getppid() != 1, not in daemon
	   if os.Getppid() != 1 {

			   cmdName := os.Args[0]
	    newProc, err := os.StartProcess(cmdName, os.Args, &os.ProcAttr{})
	    if err != nil {
		    fmt.Println("Start-Deamon: start process: %s failed, err is: %v", cmdName, err)
		    return
	    }

	    fmt.Println("Start-Deamon: run in daemon success, pid: %v", newProc.Pid)
	    return
	}

	for {
		// will write "test" to file
		fmt.Println("test")
		time.Sleep(time.Duration(3) * time.Second)
	}
}
