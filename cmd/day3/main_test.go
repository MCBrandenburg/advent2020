package main

import "testing"

func Test_run(t *testing.T) {
	ss := readInput("./example.txt")
	type args struct {
		ss     [][]string
		slopeY int
		slopeX int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				ss:     ss,
				slopeX: 3,
				slopeY: 1,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.ss, tt.args.slopeY, tt.args.slopeX); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
