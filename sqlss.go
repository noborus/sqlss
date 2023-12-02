package sqlss

import (
	"strings"
)

// SplitQueries splits a string of semicolon-separated SQL queries into individual queries.
func SplitQueries(sql string) []string {
	// Split the string on semicolons
	queries := strings.Split(sql, ";")

	// Trim whitespace from each query
	for i, query := range queries {
		queries[i] = strings.TrimSpace(query)
	}

	return queries
}
