package main

import "testing"

func Test_countHousesPart1(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{">"}, 2},
		{"test2", args{"^>v<"}, 4},
		{"test3", args{"^v^v^v^v^v"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.line); got != tt.want {
				t.Errorf("Test_countHousesPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countHousesPart2(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"^v"}, 3},
		{"test2", args{"^>v<"}, 3},
		{"test3", args{"^v^v^v^v^v"}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.line); got != tt.want {
				t.Errorf("Test_countHousesPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
