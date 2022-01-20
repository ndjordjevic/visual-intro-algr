package palindrome

import (
	"testing"
)

func Test_remFirstLast(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should rem first and last char from 'nenad'",
			args: args{
				str: "nenad",
			},
			want: "ena",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := remFirstLast(tt.args.str); got != tt.want {
				t.Errorf("remFirstLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPal(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true that 'aba' is a palindrome",
			args: args{
				str: "aba",
			},
			want: true,
		},
		{
			name: "should return true that 'a' is a palindrome",
			args: args{
				str: "a",
			},
			want: true,
		},
		{
			name: "should return true that '' is a palindrome",
			args: args{
				str: "",
			},
			want: true,
		},
		{
			name: "should return true that 'aa' is a palindrome",
			args: args{
				str: "aa",
			},
			want: true,
		},
		{
			name: "should return false that 'aggh' is a palindrome",
			args: args{
				str: "aggh",
			},
			want: false,
		},
		{
			name: "should return false that 'kaghak' is a palindrome",
			args: args{
				str: "kaghak",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPal(tt.args.str); got != tt.want {
				t.Errorf("isPal() = %v, want %v", got, tt.want)
			}
		})
	}
}
