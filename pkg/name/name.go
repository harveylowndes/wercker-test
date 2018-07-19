package name

import "log"

// Test is just a struct with a name. Interesting, right?
type Test struct {
	name string
}

// CreateName adds "harvey" to the name property on a Test strut.
// Again, very interesting...
func CreateName(t *Test) {
	log.Print("Adding \"Harvey\" as name...")
	t.name = "Harvey"
	log.Print("Added.")
}

// GetName gets the name.
func (t *Test) GetName() string {
	return t.name
}
