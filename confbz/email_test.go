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
		{"base", EmailConf{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotConf := GetEmailConf(); !reflect.DeepEqual(gotConf, tt.wantConf) {
				t.Errorf("GetEmailConf() = %v, want %v", gotConf, tt.wantConf)
			}
		})
	}
}
