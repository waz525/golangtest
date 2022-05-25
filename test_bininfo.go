package main

import (
	"flag"
	"fmt"
	"os"
	"BinInfo"
)

func main() {
	v := flag.Bool("v", false, "show bin info")
	flag.Parse()
	if *v {
		_, _ = fmt.Fprint(os.Stderr, BinInfo.StringifyMultiLine())
		os.Exit(1)
	}

	fmt.Println("my app running...")
	fmt.Println("bye...")
}
