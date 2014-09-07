package main

import "math/rand"

type id int64

func newId() id { return id(rand.Int63()) }

type mob interface {
	id() id
	name() string
	desc() string
}

type viewer interface {
	id() id
	newline()
	print(string, color)
	println(string, color)
}
