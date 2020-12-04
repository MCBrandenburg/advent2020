package main

import (
	"reflect"
	"testing"
)

func Test_createRule(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  rulePart1
		want1 string
	}{
		{
			name: "1-3 a: abcde",
			args: args{
				s: "1-3 a: abcde",
			},
			want: rulePart1{
				min: 1,
				max: 3,
				r:   'a',
			},
			want1: "abcde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := createRulePart1(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createRule() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("createRule() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_rule_IsValid(t *testing.T) {
	type fields struct {
		min int
		max int
		r   rune
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "1-3 a: abcde",
			fields: fields{
				min: 1,
				max: 3,
				r:   'a',
			},
			args: args{
				s: "abcde",
			},
			want: true,
		},
		{
			name: "1-3 b: cdefg",
			fields: fields{
				min: 1,
				max: 3,
				r:   'b',
			},
			args: args{
				s: "cdefg",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rulePart1{
				min: tt.fields.min,
				max: tt.fields.max,
				r:   tt.fields.r,
			}
			if got := r.IsValid(tt.args.s); got != tt.want {
				t.Errorf("rule.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
