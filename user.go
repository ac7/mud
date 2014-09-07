package main

import (
	"bufio"
	"io"
)

type user struct {
	*bufio.ReadWriter
	io.Closer
}

func (u *user) Println(msg string) {
	u.WriteString(msg + "\n")
	u.Flush()
}

func newUser(stream io.ReadWriteCloser) *user {
	u := &user{
		ReadWriter: bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream)),
		Closer:     stream,
	}
	return u
}
