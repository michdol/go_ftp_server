package main


import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)


func main() {
	go func() {
		listener, err := net.Listen("tcp", "0.0.0.0:8000")
		if err != nil {
			log.Fatal(err)
		}
		for {
			log.Print("Listening?")
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err) // eg. connection aborted
				continue
			}
			go handleConn(conn, 8000) // handle one connection at a time
		}
	}()

	go func() {
		listener, err := net.Listen("tcp", "0.0.0.0:8001")
		if err != nil {
			log.Fatal(err)
		}
		for {
			log.Print("Listening?")
			conn, err := listener.Accept()
			if err != nil {
				log.Print(err) // eg. connection aborted
				continue
			}
			go handleConn(conn, 8001) // handle one connection at a time
		}
	}()

	for {

	}
}

func handleConn(c net.Conn, port int) {
	defer c.Close()
	for {
		message := fmt.Sprintf("%s from port: %d\n", time.Now().Format("15:04:05"), port)
		_, err := io.WriteString(c, message)
		if err != nil {
			return // client diconnected for eg.
		}
		time.Sleep(1 * time.Second)
	}
}