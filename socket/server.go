package main

import (
	"bufio"
	"fmt"
	"flag"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintf(conn, "Input q to quit !\r\n")
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(conn.RemoteAddr()," ---> ",string(data))
		if string(data)[0] == 'q' {
				break
		}
		fmt.Fprintf(conn, string(data))
	}
}

func main() {
		var port string
		flag.StringVar(&port, "p", "3371",  "Port for Server !")
		flag.Parse()

		fmt.Println("Listen on "+port+" ... ")
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
