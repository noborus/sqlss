package sqlss_test

import (
	"fmt"

	"github.com/noborus/sqlss"
)

func ExampleSplitQueries() {
	sql := `UPDATE users SET name='Bob' WHERE id=1;SELECT * FROM users;`
	queries := sqlss.SplitQueries(sql)
	for _, query := range queries {
		fmt.Println(query)
	}
	// Output:
	// UPDATE users SET name='Bob' WHERE id=1
	// SELECT * FROM users
}

func ExampleSplitQueries_second() {
	sql := `SELECT 'O''Reilly' ; SELECT 2;`
	queries := sqlss.SplitQueries(sql)
	for _, query := range queries {
		fmt.Println(query)
	}
	// Output:
	// SELECT 'O''Reilly'
	// SELECT 2
}
