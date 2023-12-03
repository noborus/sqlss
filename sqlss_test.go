package sqlss

import (
	"reflect"
	"testing"
)

func TestSplitQueries(t *testing.T) {
	type args struct {
		sql string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestEmpty",
			args: args{
				sql: "",
			},
			want: []string{},
		},
		{
			name: "TestNormal",
			args: args{
				sql: "SELECT * FROM users; SELECT * FROM posts;",
			},
			want: []string{
				"SELECT * FROM users",
				"SELECT * FROM posts",
			},
		},
		{
			name: "testEscapedSemicolon",
			args: args{
				sql: "SELECT * FROM users WHERE name = 'John;Doe'; SELECT * FROM posts;",
			},
			want: []string{
				"SELECT * FROM users WHERE name = 'John;Doe'",
				"SELECT * FROM posts",
			},
		},
		{
			name: "testEscapedSemicolon2",
			args: args{
				sql: "SELECT * FROM users WHERE name = 'John''Doe'; SELECT * FROM posts;",
			},
			want: []string{
				"SELECT * FROM users WHERE name = 'John''Doe'",
				"SELECT * FROM posts",
			},
		},
		{
			name: "testEscapedSemicolon3",
			args: args{
				sql: "SELECT * FROM users WHERE name = 'John'';''Doe'; SELECT * FROM posts;",
			},
			want: []string{
				"SELECT * FROM users WHERE name = 'John'';''Doe'",
				"SELECT * FROM posts",
			},
		},
		{
			name: "testComment1",
			args: args{
				sql: "SELECT * FROM users WHERE name = 'John;Doe'; -- SELECT * FROM posts;",
			},
			want: []string{
				"SELECT * FROM users WHERE name = 'John;Doe'",
			},
		},
		{
			name: "testComment2",
			args: args{
				sql: "SELECT * FROM users WHERE name = 'John;Doe'; /* SELECT * FROM posts; */",
			},
			want: []string{
				"SELECT * FROM users WHERE name = 'John;Doe'",
			},
		},
		{
			name: "testComment3",
			args: args{
				sql: "SELECT /* -- */ 'a';SELECT 'b'",
			},
			want: []string{
				"SELECT /* -- */ 'a'",
				"SELECT 'b'",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitQueries(tt.args.sql)
			for i, query := range tt.want {
				if !reflect.DeepEqual(got[i], query) {
					t.Errorf("SplitQueries() = [%v], want [%v]", got[i], query)
				}
			}
		})
	}
}
