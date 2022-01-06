package day7

import (
	"reflect"
	"testing"
)

//unit test
func Test_parseInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "1,2,3", args: args{input: "1,2,3"}, want: []int{1, 2, 3}, wantErr: false},
		{name: "a,2,3", args: args{input: "a,2,3"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInput(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minMax(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name    string
		args    args
		wantMin int
		wantMax int
		wantErr bool
	}{
		{name: "1,2,3", args: args{slice: []int{1, 2, 3}}, wantMin: 1, wantMax: 3, wantErr: false},
		{name: "empty", args: args{slice: []int{}}, wantMin: 0, wantMax: 0, wantErr: true},
		{name: "100,10,8", args: args{slice: []int{100,10,8}}, wantMin: 8, wantMax: 100, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMin, gotMax, err := minMax(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("minMax() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMin != tt.wantMin {
				t.Errorf("minMax() gotMin = %v, want %v", gotMin, tt.wantMin)
			}
			if gotMax != tt.wantMax {
				t.Errorf("minMax() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
