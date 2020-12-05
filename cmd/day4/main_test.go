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
		want []passport
	}{
		{
			args: args{
				f: "./example.txt",
			},
			want: []passport{
				{
					"byr": "1937",
					"cid": "147",
					"ecl": "gry",
					"eyr": "2020",
					"hcl": "#fffffd",
					"hgt": "183cm",
					"iyr": "2017",
					"pid": "860033327",
				},
				{
					"byr": "1929",
					"cid": "350",
					"ecl": "amb",
					"eyr": "2023",
					"hcl": "#cfa07d",
					"iyr": "2013",
					"pid": "028048884",
				},
				{
					"byr": "1931",
					"ecl": "brn",
					"eyr": "2024",
					"hcl": "#ae17e1",
					"hgt": "179cm",
					"iyr": "2013",
					"pid": "760753108",
				},
				{
					"ecl": "brn",
					"eyr": "2025",
					"hcl": "#cfa07d",
					"hgt": "59in",
					"iyr": "2011",
					"pid": "166559648",
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
		pp []passport
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				pp: []passport{
					{
						"byr": "1937",
						"cid": "147",
						"ecl": "gry",
						"eyr": "2020",
						"hcl": "#fffffd",
						"hgt": "183cm",
						"iyr": "2017",
						"pid": "860033327",
					},
					{
						"byr": "1929",
						"cid": "350",
						"ecl": "amb",
						"eyr": "2023",
						"hcl": "#cfa07d",
						"iyr": "2013",
						"pid": "028048884",
					},
					{
						"byr": "1931",
						"ecl": "brn",
						"eyr": "2024",
						"hcl": "#ae17e1",
						"hgt": "179cm",
						"iyr": "2013",
						"pid": "760753108",
					},
					{
						"ecl": "brn",
						"eyr": "2025",
						"hcl": "#cfa07d",
						"hgt": "59in",
						"iyr": "2011",
						"pid": "166559648",
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPart1(tt.args.pp); got != tt.want {
				t.Errorf("findPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPart2(t *testing.T) {
	type args struct {
		pp []passport
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			want: 0,
			args: args{
				pp: []passport{
					{
						"eyr": "1972",
						"cid": "100",
						"hcl": "#18171d",
						"ecl": "amb",
						"hgt": "170",
						"pid": "186cm",
						"iyr": "2018",
						"byr": "1926",
					},
					{
						"iyr": "2019",
						"hcl": "#602927",
						"eyr": "1967",
						"hgt": "170cm",
						"ecl": "grn",
						"pid": "012533040",
						"byr": "1946",
					},
					{
						"hcl": "dab227",
						"iyr": "2012",
						"ecl": "brn",
						"hgt": "182cm",
						"pid": "021572410",
						"eyr": "2020",
						"byr": "1992",
						"cid": "277",
					},
					{
						"hgt": "59cm",
						"ecl": "zzz",
						"eyr": "2038",
						"hcl": "74454a",
						"iyr": "2023",
						"pid": "3556412378",
						"byr": "2007",
					},
				},
			},
		},
		{
			want: 4,
			args: args{
				pp: []passport{
					{
						"pid": "087499704",
						"hgt": "74in",
						"ecl": "grn",
						"iyr": "2012",
						"eyr": "2030",
						"byr": "1980",
						"hcl": "#623a2f",
					},
					{
						"eyr": "2029",
						"ecl": "blu",
						"cid": "129",
						"byr": "1989",
						"iyr": "2014",
						"pid": "896056539",
						"hcl": "#a97842",
						"hgt": "165cm",
					},
					{
						"hcl": "#888785",
						"hgt": "164cm",
						"byr": "2001",
						"iyr": "2015",
						"cid": "88",
						"pid": "545766238",
						"ecl": "hzl",
						"eyr": "2022",
					},
					{
						"iyr": "2010",
						"hgt": "158cm",
						"hcl": "#b6652a",
						"ecl": "blu",
						"byr": "1944",
						"eyr": "2021",
						"pid": "093154719",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPart2(tt.args.pp); got != tt.want {
				t.Errorf("findPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
