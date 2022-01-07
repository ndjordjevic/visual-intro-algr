package binsearch

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		list   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Should return 3",
			args{
				[]int{1, 3, 5, 6, 7, 9},
				6,
			},
			3,
		},
		{
			name: "Should return -1",
			args: args{
				list:   []int{4, 6, 8},
				target: 10,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.list, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
