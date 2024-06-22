package ini

/*
Section implements a single section of an INI, and begins with the
name of the section in square brackets (e.g.: [Colors]).
*/
type Section struct {
	name   string
	params [][2]string
}

/*
NewSection returns a newly initialized instance of *[Section].
*/
func NewSection(name string) *Section {
	return &Section{name: name, params: make([][2]string, 0)}
}

/*
Define assigns value v to var k within the receiver instance.
*/
func (r *Section) Define(k, v string) {
	var updated bool
	for i := 0; i < len(r.params) && !updated; i++ {
		if eq(k, r.params[i][0]) {
			r.params[i][1] += v
			updated = true
		}
	}

	if !updated {
		r.params = append(r.params, [2]string{k, v})
	}
}

/*
String returns the string representation of the receiver instance.
*/
func (r Section) String() string {
	if &r == nil {
		return ``
	}

	var s string
	s += `[` + r.name + `]` +
		string(rune(10)) +
		string(rune(10))

	for _, v := range r.params {
		s += v[0] + `: ` + v[1] + string(rune(10))
	}

	return s
}

/*
Name returns the section name associated with the receiver instance,
enclosed in square brackets ([]).  A zero string is returned if the
instance is nil.
*/
func (r Section) Name() string {
	if len(r.name) > 0 {
		return `[` + r.name + `]`
	}

	return ``
}

/*
Var returns the associated value of the named string variable found
within the receiver instance, alongside a Boolean value indicative of
a defined variable.

Case is not significant in the matching process.
*/
func (r Section) Var(name string) (v string, found bool) {
	for i := 0; i < len(r.params) && !found; i++ {
		if found = eq(r.params[i][0], name); found {
			v = r.params[i][1]
		}
	}

	return
}

/*
Len returns the integer length of the receiver instance in terms of the
total number of variables defined.
*/
func (r Section) Len() int {
	return len(r.params)
}

/*
Index returns the Nth string slice within the receiver instance alongside a
Boolean value indicative of a successful index call.
*/
func (r Section) Index(idx int) (v string, found bool) {
	if found = r.Len() <= idx; found {
		v = r.params[idx][1]
	}

	return
}
