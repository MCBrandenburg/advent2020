package main

import (
	"testing"
)

func Test_name(t *testing.T) {
	type args struct {
		preamble int
		input    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				preamble: 5,
				input: []int{
					35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576,
				},
			},
			want: 127,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := name(tt.args.preamble, tt.args.input); got != tt.want {
				t.Errorf("name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_name2(t *testing.T) {
	type args struct {
		preamble int
		input    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				preamble: 25,
				input: []int{
					35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576,
				},
			},
			want: 62,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := name2(tt.args.preamble, 127, tt.args.input); got != tt.want {
				t.Errorf("name2() = %v, want %v", got, tt.want)
			}
		})
	}
}
