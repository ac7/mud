package main

type monster struct {
	monstername string // really need to find a better way to avoid name collision
}

func (m *monster) name() string { return m.monstername }
func (m *monster) desc() string { return "A big scary monster." }
