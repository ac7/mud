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

func serve(world *world) {
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
		world.AddUser(newUser(conn))
	}
}

func main() {
	world := newWorld()
	shouldServe := flag.Bool("serve", false, "Act as a telnet server")
	flag.Parse()
	if *shouldServe {
		serve(world)
	} else {
		// build a mock connection out of stdin and stdout
		stream := struct {
			io.Reader
			io.Writer
			io.Closer
		}{os.Stdin, os.Stdout, os.Stdout}
		world.AddUser(newUser(stream))
	}
}
