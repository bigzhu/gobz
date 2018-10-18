package rsabz

import (
	"crypto/rsa"
	"reflect"
	"testing"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA11coKTbUugNt9ljSQHTItLryyvAH1EI4Ww+mJyrbwAf8mnLl7kPAWBGYeS5hl5SZTDB8D1NhnK8JjU5iScagwcFbn+LKT1Z/8BdLiezwbPpSEtYWQJtb1tpgxXNYwWeYyJGhnp7d3f+/rPhBG2Y36mWjNYCAESXP+KtuXdhvsjAVccp6yRJlyiTewa6nrWBT1A5O5Lfs17XOfPqKmLWxWA30rtyJT3oMdENNY/aex0n+P5xI6eUGCHBlOukXPObBfLgHujUiDgatsKZKSNPeNRUgMsL8TLt+9p8rlZpDNfb3KCWFxVe6zm4KuLcyPpvTuy3ybZ7Gk1+doFF0keSuIQIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	type args struct {
		plainText string
		signatur  string
		publicKey string
	}
	ok := args{plainText: "dfjdifjid=4&sign_type=RSA2", signatur: "rHariRxHXaDZ1Pmvy6TKXVvvCP91xpR6VBIofPnSMSl/hr2ZejCVGGW7PNOhCWgspcpzLi19JOVoAZtCHk2XKiU7OP5/FShRCWqlEI0gGa7JBLF7IcrsTMcVOq0LCUlQjBhAFZ3LIEtOVCzXTQ61A7bjrn1mWrM9uwIx63AtMZj3/pb0oUPzvFZKxi5r0R5XHOObjDwcBhI3kv7460cb9a/mMgT4sQwOJISogRNDbuvuFhTsbNyHsKICFVndikvhki9f9xtnuSwpvf5k0uj1/n2gSiRt//hiLMJGK0HIgX7uc8EAi9RfsevxMRGISqeQaGvtCn5y7a6nvMT4OGmQpg==", publicKey: publicKey}
	errArg := args{plainText: "dfjdifjid=4&sign_type=RSA2$fuck=you", signatur: "rHariRxHXaDZ1Pmvy6TKXVvvCP91xpR6VBIofPnSMSl/hr2ZejCVGGW7PNOhCWgspcpzLi19JOVoAZtCHk2XKiU7OP5/FShRCWqlEI0gGa7JBLF7IcrsTMcVOq0LCUlQjBhAFZ3LIEtOVCzXTQ61A7bjrn1mWrM9uwIx63AtMZj3/pb0oUPzvFZKxi5r0R5XHOObjDwcBhI3kv7460cb9a/mMgT4sQwOJISogRNDbuvuFhTsbNyHsKICFVndikvhki9f9xtnuSwpvf5k0uj1/n2gSiRt//hiLMJGK0HIgX7uc8EAi9RfsevxMRGISqeQaGvtCn5y7a6nvMT4OGmQpg==", publicKey: publicKey}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", ok, false},
		{"error", errArg, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Verify(tt.args.plainText, tt.args.signatur, tt.args.publicKey); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getPublicKey(t *testing.T) {
	type args struct {
		publicKey string
	}
	arg := args{publicKey: publicKey}
	tests := []struct {
		name    string
		args    args
		wantP   *rsa.PublicKey
		wantErr bool
	}{
		{"base", arg, nil, false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, err := getPublicKey(tt.args.publicKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPublicKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("getPublicKey() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}
