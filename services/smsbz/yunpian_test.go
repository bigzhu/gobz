package smsbz

import (
	"testing"
)

func TestSendIntl(t *testing.T) {
	type args struct {
		phone string
		msg   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"no country code prefix", args{"15087190161", "123456"}, true},
		{"chinese mobile", args{"+8615087190161", "123456"}, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SendIntl(tt.args.phone, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendIntl error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

	//fmt.Printf("%#v", isIntlMobile("15087190161"))
}
