package main

type room struct {
	name, desc string
	exits      map[string]*room
	contents   []mob
}

func (r *room) describe(viewer *user) {
	viewer.println("-=[{ "+r.name+" }]=-", green)
	viewer.newline()
	viewer.println(r.desc, none)
	viewer.newline()
}
