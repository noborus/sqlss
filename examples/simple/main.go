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
