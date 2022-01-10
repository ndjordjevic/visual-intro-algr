package selsort

import (
	"reflect"
	"testing"
)

func Test_selSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Should sort the array",
			args: args{
				list: []int{4, 1, 5, 3, 2},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Should sort successfuly",
			args: args{
				list: []int{4, 3, 2, 1, 0, -1, -99},
			},
			want: []int{-99, -1, 0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selSort(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
