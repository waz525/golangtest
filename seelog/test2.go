
package main

import (
    "log"
)

func init() {
    log.SetPrefix("TRACE: ")
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.SetFlags(log.Ldate | log.Ltime )
}

func main() {
    // Println writes to the standard logger.
    log.Println("message")

    // Fatalln is Println() followed by a call to os.Exit(1).
//    log.Fatalln("fatal message")

    // Panicln is Println() followed by a call to panic().
//    log.Panicln("panic message")
}
