package main

type room struct {
	name, desc string
	exits      map[string]*room
	contents   []mob
}

func (r *room) describeTo(viewer *user) {
	viewer.println("-=[{ "+r.name+" }]=-", green)
	viewer.newline()
	viewer.println(r.desc, none)
	viewer.print("Here: ", black)
	for _, thing := range r.contents {
		if thing != viewer {
			viewer.print(thing.name()+", ", black)
		}
	}
	viewer.newline()
}
