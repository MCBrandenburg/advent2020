package main

import (
	"reflect"
	"testing"
)

func Test_getCommands(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name string
		args args
		want []command
	}{
		{
			name: "example",
			args: args{
				ss: []string{
					"nop +0",
					"acc +1",
					"jmp +4",
					"acc +3",
					"jmp -3",
					"acc -99",
					"acc +1",
					"jmp -4",
					"acc +6",
				},
			},
			want: []command{
				{"nop", 0},
				{"acc", 1},
				{"jmp", 4},
				{"acc", 3},
				{"jmp", -3},
				{"acc", -99},
				{"acc", 1},
				{"jmp", -4},
				{"acc", 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCommands(tt.args.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPart1(t *testing.T) {
	type args struct {
		cc []command
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPart1(tt.args.cc); got != tt.want {
				t.Errorf("findPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_execute(t *testing.T) {
	type args struct {
		maxRun int
		cc     []command
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				maxRun: 1,
				cc: []command{
					{"nop", 0},
					{"acc", 1},
					{"jmp", 4},
					{"acc", 3},
					{"jmp", -3},
					{"acc", -99},
					{"acc", 1},
					{"jmp", -4},
					{"acc", 6},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := execute(tt.args.maxRun, tt.args.cc); got != tt.want {
				t.Errorf("execute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executeCorrectly(t *testing.T) {
	type args struct {
		cc []command
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "example",
			args: args{
				cc: []command{
					{"nop", 0},
					{"acc", 1},
					{"jmp", 4},
					{"acc", 3},
					{"jmp", -3},
					{"acc", -99},
					{"acc", 1},
					{"nop", -4},
					{"acc", 6},
				},
			},
			want:  8,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := executeCorrectly(tt.args.cc)
			if got != tt.want {
				t.Errorf("executeCorrectly() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("executeCorrectly() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_findPart2(t *testing.T) {
	type args struct {
		cc []command
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				cc: []command{
					{"nop", 0},
					{"acc", 1},
					{"jmp", 4},
					{"acc", 3},
					{"jmp", -3},
					{"acc", -99},
					{"acc", 1},
					{"jmp", -4},
					{"acc", 6},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPart2(tt.args.cc); got != tt.want {
				t.Errorf("findPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
