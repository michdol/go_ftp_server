package connection

import (
	"bufio"
	"io"
	"log"
	"net"
)


type Connection struct {
	conn			net.Conn
	reader		*bufio.Reader
	writer		*bufio.Writer
}

func (connection Connection) Serve() {
	log.Print("Connection is serving")
	for {
		line, err := connection.reader.ReadString('\n')
		if err != nil {
			log.Print(err == io.EOF)
			log.Fatal(err)
			break
		}
		connection.receiveLine(line)
	}
	connection.conn.Close()
	log.Print("Connection terminated")
}

func (connection *Connection) receiveLine(line string) {
	log.Print("Received line: ", line)
}

func NewConnection(conn net.Conn) *Connection {
	c := new(Connection)
	c.conn = conn
	c.reader = bufio.NewReader(conn)
	c.writer = bufio.NewWriter(conn)
	return c
}