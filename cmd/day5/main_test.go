package main

import "testing"

func Test_calcSeatID(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "FBFBBFFRLR",
			args: args{
				s: "FBFBBFFRLR",
			},
			want: 357,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcSeatID(tt.args.s); got != tt.want {
				t.Errorf("calcSeatID() = %v, want %v", got, tt.want)
			}
		})
	}
}
