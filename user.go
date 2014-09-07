package main

import (
	"bufio"
	"io"
	"net"
)

type user struct {
	*bufio.ReadWriter
	io.Closer
}

func (u *user) Println(msg string) {
	u.WriteString(msg + "\n")
	u.Flush()
}

func newUser(conn net.Conn) *user {
	u := &user{
		ReadWriter: bufio.NewReadWriter(
			bufio.NewReader(conn), bufio.NewWriter(conn)),
		Closer: conn,
	}
	return u
}
