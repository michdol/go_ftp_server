package main

import (
	"log"

	"ftpserver/server"
)


func main() {
	s := server.NewServer()
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
