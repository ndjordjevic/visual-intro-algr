package swap

import (
	"reflect"
	"testing"
)

func Test_swap(t *testing.T) {
	type args struct {
		list  []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Should swap the two values at given indexes",
			args: args{
				list:  []int{4, 7, 4, 8, 10},
				start: 1,
				end:   3,
			},
			want: []int{4, 8, 4, 7, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swap(tt.args.list, tt.args.start, tt.args.end)
			if !reflect.DeepEqual(tt.args.list, tt.want) {
				t.Errorf("selSort() = %v, want %v", tt.args.list, tt.want)
			}
		})
	}
}
