package main

import (
	"bufio"
	"io"
)

type user struct {
	*bufio.ReadWriter // for WriteString and ReadString
	io.Closer
	_id id

	_name string
	room  *room
}

func (u *user) id() id       { return u._id }
func (u *user) name() string { return u._name }
func (u *user) desc() string { return "[ " + u._name + " ]" }
func (u *user) look()        { u.room.describeTo(u) }
func (u *user) newline()     { u.println(``, none) }

func (u *user) print(msg string, c color) {
	u.WriteString(string(c) + msg + string(resetCode))
	u.Flush()
}

func (u *user) println(msg string, c color) {
	u.WriteString(string(c) + msg + string(resetCode) + "\n")
	u.Flush()
}

func (u *user) disconnect() {
	u.println(`... disconnecting`, red)
	u.newline()
	u.Close()
}

func (u *user) moveInto(room *room) {
	if u.room != nil {
		// leave handler here later
	}
	u.room = room
	room.contents = append(room.contents, u)
	u.println("You move into "+u.room.name, blue)
}

func newUser(stream io.ReadWriteCloser) *user {
	u := &user{
		ReadWriter: bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream)),
		Closer:     stream,
		_id:        newId(),
	}
	return u
}
