package httpbz

import "testing"

func TestAddParamsToUrl(t *testing.T) {
	type args struct {
		old    string
		params map[string]string
	}
	p := map[string]string{
		"id":    "1",
		"index": "10",
	}
	arg := args{old: "https://bigzhu.net", params: p}
	tests := []struct {
		name    string
		args    args
		wantNew string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"base", arg, "https://bigzhu.net?id=1&index=10", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNew, err := AddParamsToUrl(tt.args.old, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddParamsToUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNew != tt.wantNew {
				t.Errorf("AddParamsToUrl() = %v, want %v", gotNew, tt.wantNew)
			}
		})
	}
}
