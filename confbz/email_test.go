package confbz

import (
	"reflect"
	"testing"
)

func TestGetEmailConf(t *testing.T) {
	tests := []struct {
		name     string
		wantConf EmailConf
	}{
		// TODO: Add test cases.
		{"base", EmailConf{User: "oeohomos@163.com", Password: "abc123456", Host: "smtp.163.com", Port: 25}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotConf := GetEmailConf(); !reflect.DeepEqual(gotConf, tt.wantConf) {
				t.Errorf("GetEmailConf() = %v, want %v", gotConf, tt.wantConf)
			}
		})
	}
}
