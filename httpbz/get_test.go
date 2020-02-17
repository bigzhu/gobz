package httpbz

import (
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	type args struct {
		url    string
		client *http.Client
	}
	arg1 := args{url: "https://follow.bigzhu.net/api/OauthInfo", client: nil}
	arg2 := args{url: "https://follow.bigzhu.net/api/OldMessages?before=1911-08-28T15:42:45%2B08:00", client: nil}
	tests := []struct {
		name           string
		args           args
		wantData       string
		wantStatusCode int
		wantErr        bool
	}{
		// TODO: Add test cases.
		{"401", arg1, `{"error":"need login"}`, 401, false},
		{"200", arg2, `[]`, 200, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, gotStatusCode, err := Get(tt.args.url, tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotData != tt.wantData {
				t.Errorf("Get() gotData = %v, want %v", gotData, tt.wantData)
			}
			if gotStatusCode != tt.wantStatusCode {
				t.Errorf("Get() gotStatusCode = %v, want %v", gotStatusCode, tt.wantStatusCode)
			}
		})
	}
}
