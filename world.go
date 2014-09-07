package main

type world struct {
	startZone *room
}

func (w *world) AddUser(u *user) {
	u.newline()
	u.println(`... connected`, red)
	u.moveInto(w.startZone)
	u.look()

	u.disconnect()
}

func newWorld() *world {
	return &world{
		startZone: &room{
			name:  "The Start Zone",
			desc:  "This is the starting zone. You can walk around here if you feel like it.",
			exits: map[string]*room{},
			contents: []mob{
				&monster{"grue", newID()},
			},
		},
	}
}
