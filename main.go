package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

const port = 1030

func serve() {
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

func main() {
	shouldServe := flag.Bool("serve", false, "Act as a telnet server")
	flag.Parse()
	if *shouldServe {
		serve()
	} else {
		// build a mock connection out of stdin and stdout
		stream := struct {
			io.Reader
			io.Writer
			io.Closer
		}{os.Stdin, os.Stdout, os.Stdout}

		user := newUser(stream)
		user.Println(`Welcome 2`)
		user.Close()
	}
}
