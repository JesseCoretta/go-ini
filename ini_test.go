package ini

import (
	"bytes"
	"testing"
)

func TestSection_Parse(t *testing.T) {
	n := New()

	// test a cumulative parse involving multiple
	// []byte instances, possibly acquired from
	// multiple files.
	for _, b := range [][]byte{
		testINI,
		testINI2,
	} {
		if err := n.Parse(bytes.NewBuffer(b)); err != nil {
			t.Errorf("%s failed: %v", t.Name(), err)
			return
		}
	}

	t.Logf("GLOBAL: %s\n", n.Section(`globaL`))

	for i := 0; i < n.Len(); i++ {
		t.Logf("%s\n", n.Index(i).String())
	}

	popped := n.Pop()
	t.Logf("POPPED :: %s\n", popped)

	t.Logf("FULL: %s\n", n)
}

var testINI []byte = []byte(`[global]

; this is a comment
var1: this is a variable
var2: this is
  a multiline
  variable

var3: true

[subsection1]

var1: this is a variable too
var2: Jerry.
   Hello.`)

var testINI2 []byte = []byte(`[subsection2]

; this is a comment

var1: this is data
var2: ...
`)
