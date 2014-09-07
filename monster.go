package main

type monster struct {
	_name string
	_id   id
}

func (m *monster) id() id       { return m._id }
func (m *monster) name() string { return m._name }
func (m *monster) desc() string { return "A big scary monster." }
