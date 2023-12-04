package sqlss

import (
	"strings"
)

// SplitQueries splits a string of semicolon-separated SQL queries into individual queries.
// It does not split semicolons that are within single quotes, double quotes, back quotes or escaped by two single quotes.
func SplitQueries(sql string) []string {
	var queries []string
	var currentQuery strings.Builder
	inSingle, inDouble, inBack, escaped, inComment, inComment2 := false, false, false, false, false, false
	skip := false

	for i, r := range sql {
		if skip {
			skip = false
			currentQuery.WriteRune(r)
			continue
		}
		var next byte
		if i < len(sql)-1 {
			next = sql[i+1]
		}
		switch r {
		case '\'':
			if inSingle && next == '\'' {
				skip = true
				escaped = !escaped
			} else {
				inSingle = !inSingle
			}
			currentQuery.WriteRune(r)
		case '"':
			if !escaped {
				inDouble = !inDouble
			}
			currentQuery.WriteRune(r)
		case '`':
			if !escaped {
				inBack = !inBack
			}
			currentQuery.WriteRune(r)
		case '-':
			if next == '-' && !inSingle && !inDouble && !inBack && !escaped && !inComment2 {
				skip = true
				inComment = true
			}
			currentQuery.WriteRune(r)
		case '\n':
			inComment = false
			currentQuery.WriteRune(r)
		case '/':
			if next == '*' {
				skip = true
				inComment2 = true
			}
			currentQuery.WriteRune(r)
		case '*':
			if next == '/' {
				skip = true
				inComment2 = false
			}
			currentQuery.WriteRune(r)
		case ';':
			if inSingle || inDouble || inBack || escaped || inComment || inComment2 {
				currentQuery.WriteRune(r)
			} else {
				queries = append(queries, strings.TrimSpace(currentQuery.String()))
				currentQuery.Reset()
			}
			escaped = false
		default:
			escaped = false
			currentQuery.WriteRune(r)
		}
	}

	// Add the last query if it's not empty
	lastQuery := strings.TrimSpace(currentQuery.String())
	if lastQuery != "" {
		queries = append(queries, lastQuery)
	}

	return queries
}
