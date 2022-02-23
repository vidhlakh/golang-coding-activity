package countbits

import "testing"

func Test_countbits(t *testing.T) {
	type args struct {
		input int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "right testcase1",
			args: args{
				input: 6,
			},
			want: 2,
		},
		{
			name: "testcase2",
			args: args{
				input: 13,
			},
			want: 3,
		},
		{
			name: "testcase3",
			args: args{
				input: 0,
			},
			want: 0,
		},
		{
			name: "testcase4",
			args: args{
				input: -4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountbitsMethod2(tt.args.input); got != tt.want {
				t.Errorf("countbits() = %v, want %v", got, tt.want)
			}
		})
	}
}
