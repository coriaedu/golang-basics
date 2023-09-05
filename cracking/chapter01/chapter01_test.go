package chapter01

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{
			name: "Empty string",
			word: "",
			want: true,
		},
		{
			name: "Single character",
			word: "s",
			want: true,
		},
		{
			name: "Same character repeated",
			word: "ss",
			want: false,
		},
		{
			name: "All Unique",
			word: "asdfgh",
			want: true,
		},
		{
			name: "First and last char are the same",
			word: "abcdefghijka",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUnique(tt.word); got != tt.want {
				t.Errorf("IsUnique(%v) = %v, want %v", tt.word, got, tt.want)
			}
		})
	}
}

func TestCheckPermutation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Empty strings",
			args: args{"", ""},
			want: true,
		},
		{
			name: "One empty string",
			args: args{"notempty", ""},
			want: false,
		},
		{
			name: "Another empty string",
			args: args{"", "something"},
			want: false,
		},
		{
			name: "Happy case",
			args: args{"happy", "yapph"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPermutation(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CheckPermutation(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
