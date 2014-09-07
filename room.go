package main

type room struct {
	name, desc string
	exits      map[string]*room
	contents   []mob
}

func (r *room) describeTo(viewer viewer) {
	viewer.newline()
	viewer.println("-=[{ "+r.name+" }]=-", green)
	viewer.println(" "+r.desc, none)
	viewer.print(" Here: ", black)
	for _, thing := range r.contents {
		if thing.id() != viewer.id() {
			viewer.print(thing.name()+", ", black)
		}
	}
	viewer.newline()
	viewer.newline()
}
