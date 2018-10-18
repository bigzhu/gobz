package rsabz

import (
	"crypto/rsa"
	"reflect"
	"testing"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA11coKTbUugNt9ljSQHTItLryyvAH1EI4Ww+mJyrbwAf8mnLl7kPAWBGYeS5hl5SZTDB8D1NhnK8JjU5iScagwcFbn+LKT1Z/8BdLiezwbPpSEtYWQJtb1tpgxXNYwWeYyJGhnp7d3f+/rPhBG2Y36mWjNYCAESXP+KtuXdhvsjAVccp6yRJlyiTewa6nrWBT1A5O5Lfs17XOfPqKmLWxWA30rtyJT3oMdENNY/aex0n+P5xI6eUGCHBlOukXPObBfLgHujUiDgatsKZKSNPeNRUgMsL8TLt+9p8rlZpDNfb3KCWFxVe6zm4KuLcyPpvTuy3ybZ7Gk1+doFF0keSuIQIDAQABAoIBABT+LS8fZkugC/dJH24jNHHSC9tW3RyHPCE8fFEzL2sgHDMeYmYA+t5bEMc3i1IQrYUs8Ns5oLl6ou7ApnpUv9VDuV/z5w8m+p/4VWfjzmiBWsD1WAQCYUinz2lQXmW0aMOc1um0ySJcTV5acVZvd6IFOIz9sv549jZi2AaUFAeF1Lni12Mst9uCrU1bhj4jZ6XNfOMzJK26SJkCaaDb5FjgYDL0Y40cqVhQ7wv7W7wpmc4E3YgRcPxXffT82yKtyMnMuzs6Yw9CFhyGWctlCyMohhA6hDzTvcyu506r4j3cpw/gdUbJO7SQnXWBovYDQeqtbqbZjG5B5ayfpA4VR1UCgYEA65LXZb6jIKPd3W0xYYbceYbo+t6lW+kd+WiBicFbarVdwY6TjGQK4TuL0jKICrE1fXr065Ox1HIGNB3Reqg6OBJyLmoxbHWDNbYku2H+XmHUAMX3h65g/gbFbPKePISScXXIq8LIQQtcaRZnpLw8EdiJV4nfj75YhATCMFtGcjcCgYEA6gMwaqrTeyc3tNZRxnboRh5MxLW7s6MGVElxfB9LpHxvAS/jbuVBT84plTN/0cV89r8EbCo0mA1mdBbt2t3NiOkwsPb/K4hcJXPGqgncVGwrt5OO5I6lulat3qDu7Vbr4MCGt6lEPyY5+tKMwl/JLm87DOoR3XI+raQpbJ0kFmcCgYBW30fflXXCy6s9xM65EGqtxjqUiRah1k3fc9KRYmS77TYt8s4mr53zgWoGkevpRYBcTcOUPFsP5Ry0i2p8uRz/II5K345qWMuSOyFAmNqdI6Zywo6N0piiFQ0pKC+CUcFPZggijyoaEn55onGGJoRgcGH1jjqLX750sOyLuJ058QKBgFAQzUUBs5/kWCv4VDxaPLiY710ybMlWkvjNtBVqw7/Aj15Rmzg7GjCK8jFJUIdeUZY+3u1/N0V6+D3qBnGnCAa6+lzh9YZYfmaWCTfubcM4y3Bg/Tf4En4B4DVWzMjJv6TDWpSGRoJwus+1sFFk4nmPtx2BDKG+z/ErRL0TuuArAoGAHMtG1traJVI4cjUiQzh19NpWcUgcsNQyYM+Qtii7wTLLJMeqlIucaXRbtX8lTATeI4DWvUKPnlL5KBINTZVNcL5vZE/AdZm8VqA+BF6M3w3V5AjtfXSXpgRIAHOq8QpZsALnff7s54PpMvbQIqjAYU/y/qViS9078lHG2a2myTE=
-----END RSA PRIVATE KEY-----`

func TestSign(t *testing.T) {
	type args struct {
		plainText  string
		privateKey string
	}

	arg := args{plainText: "dfjdifjid=4&sign_type=RSA2", privateKey: privateKey}
	tests := []struct {
		name         string
		args         args
		wantSignatur string
		wantErr      bool
	}{
		{"base", arg, "rHariRxHXaDZ1Pmvy6TKXVvvCP91xpR6VBIofPnSMSl/hr2ZejCVGGW7PNOhCWgspcpzLi19JOVoAZtCHk2XKiU7OP5/FShRCWqlEI0gGa7JBLF7IcrsTMcVOq0LCUlQjBhAFZ3LIEtOVCzXTQ61A7bjrn1mWrM9uwIx63AtMZj3/pb0oUPzvFZKxi5r0R5XHOObjDwcBhI3kv7460cb9a/mMgT4sQwOJISogRNDbuvuFhTsbNyHsKICFVndikvhki9f9xtnuSwpvf5k0uj1/n2gSiRt//hiLMJGK0HIgX7uc8EAi9RfsevxMRGISqeQaGvtCn5y7a6nvMT4OGmQpg==", false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSignatur, err := Sign(tt.args.plainText, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSignatur != tt.wantSignatur {
				t.Errorf("Sign() = %v, want %v", gotSignatur, tt.wantSignatur)
			}
			t.Log(gotSignatur)
		})
	}
}

func Test_getPrivateKey(t *testing.T) {
	type args struct {
		privateKey string
	}
	arg := args{privateKey: privateKey}
	tests := []struct {
		name    string
		args    args
		want    *rsa.PrivateKey
		wantErr bool
	}{
		{"base", arg, nil, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPrivateKey(tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPrivateKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
