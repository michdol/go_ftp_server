package main


import (
	"fmt"
	"bufio"
	"log"
	"net"
	"strings"
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
			go handleConn(conn) // handle one connection at a time
		}
	}()
	for {
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}