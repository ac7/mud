package main

import (
	"log"
	"net"
	"strconv"
)

const port = 1030

func main() {
	ln, err := net.Listen(`tcp`, `:`+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Printf(`Open on localhost:%d`, port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		user := newUser(conn)
		user.Println(`Welcome`)
		user.Close()
	}
}
