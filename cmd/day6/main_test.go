package main

import (
	"reflect"
	"testing"
)

func Test_readInput(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want []group
	}{
		{
			name: "example.txt",
			args: args{
				f: "./example.txt",
			},
			want: []group{
				{
					0:  1,
					97: 1,
					98: 1,
					99: 1,
				},
				{
					0:  3,
					97: 1,
					98: 1,
					99: 1,
				},
				{
					0:  2,
					97: 2,
					98: 1,
					99: 1,
				},
				{
					0:  4,
					97: 4,
				},
				{
					0:  1,
					98: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readInput(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPart1(t *testing.T) {
	type args struct {
		gg []group
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				gg: []group{
					{
						0:  1,
						97: 1,
						98: 1,
						99: 1,
					},
					{
						0:  3,
						97: 1,
						98: 1,
						99: 1,
					},
					{
						0:  2,
						97: 2,
						98: 1,
						99: 1,
					},
					{
						0:  4,
						97: 4,
					},
					{
						0:  1,
						98: 1,
					},
				},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPart1(tt.args.gg); got != tt.want {
				t.Errorf("findPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPart2(t *testing.T) {
	type args struct {
		gg []group
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				gg: []group{
					{
						0:  1,
						97: 1,
						98: 1,
						99: 1,
					},
					{
						0:  3,
						97: 1,
						98: 1,
						99: 1,
					},
					{
						0:  2,
						97: 2,
						98: 1,
						99: 1,
					},
					{
						0:  4,
						97: 4,
					},
					{
						0:  1,
						98: 1,
					},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPart2(tt.args.gg); got != tt.want {
				t.Errorf("findPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
