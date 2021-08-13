package internal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_parentPath(t *testing.T) {
	t.Parallel()

	type args struct {
		depth int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple",
			args: args{depth: 0},
			want: ".",
		}, {
			name: "one dir up",
			args: args{depth: 1},
			want: "..",
		},
		{
			name: "two dirs up",
			args: args{depth: 2},
			want: "../..",
		},
		{
			name: "three dirs up",
			args: args{depth: 3},
			want: "../../..",
		},
		{
			name: "negative depth",
			args: args{depth: -1},
			want: ".",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, parentPath(tt.args.depth))
		})
	}
}
