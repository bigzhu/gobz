package functionbz

import (
	"log"
	"testing"
)

func Test_getRunPath(t *testing.T) {
	log.Println()
	tests := []struct {
		name     string
		wantPath string
		wantErr  bool
	}{
		{"base", "/Users/bigzhu/go/src/github.com/bigzhu/gobz/functionbz", false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPath, err := GetRunPath()
			if (err != nil) != tt.wantErr {
				t.Errorf("getRunPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPath != tt.wantPath {
				t.Errorf("getRunPath() = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}
