package minsubarr

import "testing"

func Test_indexOfMinimum(t *testing.T) {
	type args struct {
		list   []int
		startI int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should return 4",
			args: args{
				[]int{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,13,59,61,67,71,73,79,83,89,97},
				5,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := indexOfMinimum(tt.args.list, tt.args.startI); got != tt.want {
				t.Errorf("indexOfMinimum() = %v, want %v", got, tt.want)
			}
		})
	}
}
