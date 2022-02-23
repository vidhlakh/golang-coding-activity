package reversewords

import "testing"

func Test_reverseusingstack(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				input: "i like to eat ice cream",
			},
			want: "cream ice eat to like i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseusingstack(tt.args.input)
			if tt.want != got {
				t.Errorf("Expected : %v, Got: %v", tt.want, got)
			}
		})

	}
}
