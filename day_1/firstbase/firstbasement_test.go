package main

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Balanced parentheses", args: args{str: "(())"}, want: 0},
		{name: "More opening parentheses", args: args{str: "((())"}, want: 0},
		{name: "More closing parentheses", args: args{str: "(()))"}, want: 5},
		{name: "Empty string", args: args{str: ""}, want: 0},
		{name: "No parentheses", args: args{str: "abc"}, want: 0},
		{name: "Nested parentheses", args: args{str: "((()))"}, want: 0},
		{name: "Unmatched opening and closing", args: args{str: "(()())())"}, want: 9},
		{name: "Unmatched opening", args: args{str: "(()()("}, want: 0},
		{name: "Unmatched closing", args: args{str: "())()()"}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.str); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
