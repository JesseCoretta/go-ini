package ini

import (
	"bufio"
	"io"
)

/*
INI implements an INI-style configuration structure comprised of one or
more instances of [Section].
*/
type INI []*Section

/*
New initializes and returns a new instance of [INI].
*/
func New() INI {
	return make([]*Section, 0)
}

/*
Section returns an instance of *[Section] bearing the input name value.
A nil return instance indicates no matching *[Section] was found.

Case is not significant in the matching process, and encapsulating square
brackets ([]) need not be specified.
*/
func (r INI) Section(name string) *Section {
	name = trim(name, `[]`)
	for i := 0; i < len(r); i++ {
		s := r[i]
		if eq(s.name, name) {
			return s
		}
	}

	return nil
}

/*
String returns the string representation of the receiver instance.
*/
func (r INI) String() string {
	var s string
	for _, v := range r {
		if v != nil {
			s += v.String() + string(rune(10))
		}
	}

	return s
}

/*
Len returns the integer length of the receiver instance, representing
the number of *[Section] instances residing within.
*/
func (r INI) Len() int {
	return len(r)
}

/*
Index returns the Nth instance of *[Section] found within the receiver
instance.
*/
func (r INI) Index(idx int) *Section {
	if idx <= r.Len() {
		return r[idx]
	}

	return nil
}

/*
Push appends an instance of *[Section] to the receiver instance in LIFO
context.
*/
func (r *INI) Push(x *Section) {
	if x != nil {
		*r = append(*r, x)
	}
}

/*
Pop returns (and removes) the final slice of the receiver instance in
LIFO context.
*/
func (r *INI) Pop() *Section {
	if l := r.Len(); l > 0 {
		x := r.Index(l - 1)
		(*r) = (*r)[:l-1]
		return x
	}

	return nil
}

/*
Parse returns an error following an attempt to read the input [io.Reader]
instance into an instance of *[INI].
*/
func (r *INI) Parse(reader io.Reader) (err error) {
	scanner := bufio.NewScanner(reader)
	var name string

	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := trimS(line)

		if trimmedLine == "" || hasPfx(trimmedLine, ";") {
			continue
		}

		if hasPfx(trimmedLine, "[") && hasSfx(trimmedLine, "]") {
			name = trimmedLine[1 : len(trimmedLine)-1]
			*r = append(*r, NewSection(name))
			continue
		}

		idx := len(*r) - 1
		if idx < 0 {
			idx = 0
		}

		if (*r)[idx] == nil {
			(*r)[idx] = NewSection(name)
		}

		pidx := len((*r)[idx].params) - 1

		if line[0] == ' ' || line[0] == '\t' {
			// Previous line value continued
			(*r)[idx].params[pidx][1] += string(rune(32)) + trimmedLine
		} else {
			parts := splitn(trimmedLine, ":", 2)
			if len(parts) == 2 {
				key := trimS(parts[0])
				value := trimS(parts[1])
				(*r)[idx].params = append((*r)[idx].params, [2]string{key, value})
			}
		}
	}

	return
}
