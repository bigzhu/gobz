package confbz

import "testing"

func Test_getRunPath(t *testing.T) {
	tests := []struct {
		name     string
		wantPath string
		wantErr  bool
	}{
		{"base", "", false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPath, err := getRunPath()
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
