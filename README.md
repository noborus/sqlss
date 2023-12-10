# sqlss

[![PkgGoDev](https://pkg.go.dev/badge/github.com/noborus/sqlss)](https://pkg.go.dev/github.com/noborus/sqlss)

Split SQL into statements.

Do not split within single quotes, double quotes, or back quotes.
It also doesn't break at semicolons in comments.

## Usage

```go
package main

import (
	"fmt"

	"github.com/noborus/sqlss"
)

func main() {
	queries := sqlss.SplitQueries("SELECT * FROM table1;SELECT * FROM table2;")
	for _, query := range queries {
		fmt.Println(query)
	}
}

// Output:
// SELECT * FROM table1
// SELECT * FROM table2
```
