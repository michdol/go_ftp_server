package server

import (
	"log"
	"net"

	"ftpserver/connection"
)


type Server struct {
	listenTo	string

	listener	net.Listener
}

func NewServer() *Server {
	s := new(Server)
	s.listenTo = "localhost:8000"
	return s
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.listenTo)
	if err != nil {
		log.Fatal(err)
		return err
	}

	s.listener = listener

	for {
		log.Print("Listening for connections")
		conn, err := s.listener.Accept()
		if err != nil {
			log.Print("=== Connection initialization failure ===")
			log.Print(err)
			log.Print("=== Connection initialization failure ===")
			continue
		} else {
			ftpConn := connection.NewConnection(conn)
			go ftpConn.Serve()
		}
	}
}
