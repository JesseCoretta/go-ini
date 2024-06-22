# INI

Basic INI syntax parsing and abstraction

## License

The go-ini package is released under the terms of the MIT license.  For
full details, see the LICENSE file in the package root.

## Syntax Example

The following is a valid INI example. Parsing the following data would
produce an `INI` instance comprised of two (2) \*`Section` slices.

```
[sectionname]

var1: value
var2: multi
  line
  values
  are
  cool

[subsectionname]

var1: coolness
whatever: ok
```

## Example

```
package main

import (
	"fmt"
	"os"

	"github.com/JesseCoretta/go-ini"
)

func main() {

	file, err := os.Open("config.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	i := New()
	if err := i.Parse(file); err != nil {
		fmt.Println(err)
		return
	}

	sec := i.Index(0) // get 0th index (section)
	val, found := sec.Var(`var`)
	if found {
		fmt.Println("Found " + val)
		// Output: my_value
	}

	// Parse another ini into the same instance
	// while preserving the original content
        file, err = os.Open("config2.ini")
        if err != nil {
                fmt.Println(err)
                return
        }
        defer file.Close()

        if err := i.Parse(file); err != nil {
                fmt.Println(err)
                return
        }

	// call the section labeled "global"
	sec = i.Section(`global`)
	fmt.Println(sec)
}
```
