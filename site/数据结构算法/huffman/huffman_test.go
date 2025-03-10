package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_bytesToBits(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 int
	}{
		// TODO: Add test cases.
		{"first case", args{[]byte{'1', '1', '0', '0', '1', '1', '0', '0'}}, []byte{204}, 0},
		{"xxxx case", args{[]byte{'1', '0', '1', '0'}}, []byte{160}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := bytesToBits(tt.args.bytes)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bytesToBits() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("bytesToBits() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_makefreqTable(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want [256]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makefreqTable(tt.args.filePath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makefreqTable() = %v, want %v", got, tt.want)
			}
		})
	}
	ft := makefreqTable("a.txt")
	fmt.Printf("ft: %v\n", ft)
}
