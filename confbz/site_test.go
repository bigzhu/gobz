package confbz

import (
	"reflect"
	"testing"
)

func TestGetSiteConf(t *testing.T) {
	tests := []struct {
		name         string
		wantSiteConf SiteConf
	}{
		{"base", SiteConf{Secret: "secret just secret", CookieSalt: "salt just salt", Port: 4004}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSiteConf := GetSiteConf(); !reflect.DeepEqual(gotSiteConf, tt.wantSiteConf) {
				t.Errorf("GetSiteConf() = %v, want %v", gotSiteConf, tt.wantSiteConf)
			}
		})
	}
}
