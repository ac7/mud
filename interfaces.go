package main

type mob interface {
	name() string
	desc() string
}

type describer interface {
	describe(viewer *user)
}
