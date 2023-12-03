package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/noborus/sqlss"
)

func main() {
	args := os.Args[1:]
	queryStr := strings.Join(args, " ")
	queries := sqlss.SplitQueries(queryStr)
	for _, query := range queries {
		fmt.Println(query)
	}

}
