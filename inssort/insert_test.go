package inssort

import (
	"reflect"
	"testing"
)

func Test_insertInSubArr(t *testing.T) {
	type args struct {
		array []int
		right int
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should insert 9 in the proper position",
			args: args{
				array: []int{2, 3, 5, 7, 11, 13, 9, 6},
				right: 5,
				value: 9,
			},
			want: []int{2, 3, 5, 7, 9, 11, 13, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertInSubArr(tt.args.array, tt.args.right, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_remEl(t *testing.T) {
	type args struct {
		array []int
		val   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should remove 9 from an array",
			args: args{
				array: []int{2, 3, 5, 7, 11, 13, 9, 6, 10, 56},
				val:   9,
			},
			want: []int{2, 3, 5, 7, 11, 13, 6, 10, 56},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remEl(tt.args.array, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remEl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insInArr(t *testing.T) {
	type args struct {
		array []int
		index int
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "insert 9 at position 4 ",
			args: args{
				array: []int{8, 10, 13, 18, 22, 3},
				index: 4,
				value: 9,
			},
			want: []int{8, 10, 13, 18, 9, 22, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insInArr(tt.args.array, tt.args.index, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insInArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should sort the given array",
			args: args{
				array: []int{7, 3, 18, 6, 22, 1},
			},
			want: []int{1, 3, 6, 7, 18, 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertSort2(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should sort the given array",
			args: args{
				array: []int{7, 3, 18, 6, 22, 1},
			},
			want: []int{1, 3, 6, 7, 18, 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertSort2(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertSort2() = %v, want %v", got, tt.want)
			}
		})
	}
}
