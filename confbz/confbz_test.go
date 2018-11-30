package confbz

import (
	"testing"
)

func TestGetConf(t *testing.T) {
	type args struct {
		confbzName string
		v          interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"db", args{"db", &DBConf{}}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetConf(tt.args.confbzName, tt.args.v)
		})
	}
}

func Test_getConfFileNameByStruct(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"db", args{&DBConf{}}, "db"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConfFileNameByStruct(tt.args.v); got != tt.want {
				t.Errorf("getConfFileNameByStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
