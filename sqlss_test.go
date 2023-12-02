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
			name: "SplitQueries",
			args: args{
				sql: "SELECT * FROM users; SELECT * FROM posts;",
			},
			want: []string{
				"SELECT * FROM users",
				"SELECT * FROM posts",
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
