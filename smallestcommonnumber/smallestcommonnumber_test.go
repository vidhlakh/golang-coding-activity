package smallestcommonnumber

import "testing"

func Test_smallestcommonnumber(t *testing.T) {
	type args struct {
		a []int
		b []int
		c []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				a: []int{1, 4, 12, 34, 19, 28, 34, 9},
				b: []int{23, 45, 6, 7, 8, 9, 34},
				c: []int{3, 5, 7, 9, 34, 19},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestcommonnumber(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("smallestcommonnumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
