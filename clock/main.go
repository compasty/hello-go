package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8234")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err) // e.g. connection aborted
			continue
		}
		handleConn(conn)
	}

}
