package fact

import (
	"testing"
)

func Test_fact(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should compute fact of 9",
			args: args{
				n: 9,
			},
			want: 362880,
		},
		{
			name: "should compute fact of 0",
			args: args{
				n: 0,
			},
			want: 1,
		},
		{
			name: "should compute fact of 1",
			args: args{
				n: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fact(tt.args.n); got != tt.want {
				t.Errorf("fact() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_factRec(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should return fact of 5",
			args: args{
				n: 5,
			},
			want: 120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factRec(tt.args.n); got != tt.want {
				t.Errorf("factRec() = %v, want %v", got, tt.want)
			}
		})
	}
}
