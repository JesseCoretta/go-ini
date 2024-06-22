# INI

Basic INI syntax parsing and abstraction

## License

The go-ini package is released under the terms of the MIT license.  For
full details, see the LICENSE file in the package root.

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
}

```
