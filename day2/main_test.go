package main

import "testing"

func Test_getAreaFromDimension(t *testing.T) {
	type args struct {
		dimension string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{dimension: "2x3x4"}, 58},
		{"test2", args{dimension: "1x1x10"}, 43},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAreaFromDimension(tt.args.dimension); got != tt.want {
				t.Errorf("getAreaFromDimension() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRibbonLengthFromDimension(t *testing.T) {
	type args struct {
		dimension string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{dimension: "2x3x4"}, 34},
		{"test2", args{dimension: "1x1x10"}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRibbonLengthFromDimension(tt.args.dimension); got != tt.want {
				t.Errorf("getRibbonLengthFromDimension() = %v, want %v", got, tt.want)
			}
		})
	}
}
